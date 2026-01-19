package data

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v7"
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/config"
)

// Cache Redis 缓存客户端封装
// 提供配置热重载和优雅关闭支持
type Cache struct {
	conn   *redis.Client
	config *config.CacheConfig

	mu        sync.RWMutex // 读写锁，保护 conn 字段
	cancel    context.CancelFunc
	watchDone chan struct{} // 用于通知 watch 协程退出
}

// NewCache 创建缓存实例并初始化连接
// 启动配置变更监听协程
func NewCache(cfg *config.CacheConfig) *Cache {
	cache := &Cache{
		config:    cfg,
		watchDone: make(chan struct{}),
	}

	// 初始化连接
	if err := cache.initConnection(); err != nil {
		log.Printf("failed to init redis connection: %v", err)
		// 返回 degraded cache 实例，连接为 nil
		return cache
	}

	// 启动配置变更监听
	go cache.watchConfigChange()

	return cache
}

// initConnection 初始化 Redis 连接
// 使用当前配置创建新的连接
func (c *Cache) initConnection() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := c.config

	// 关闭旧连接
	if c.conn != nil {
		c.conn.Close()
	}

	conn := redis.NewClient(&redis.Options{
		Addr:     item.Host + ":" + strconv.Itoa(item.Port),
		Password: item.Auth,
		DB:       item.DB,
	})

	// 验证连接
	if err := conn.Ping().Err(); err != nil {
		log.Printf("failed to connect to redis: %v", err)
		return err
	}

	c.conn = conn
	log.Printf("redis connected successfully: %s", conn.Options().Addr)
	return nil
}

// watchConfigChange 监听配置变更并重新初始化连接
// 支持优雅关闭，通过 channel 信号控制
func (c *Cache) watchConfigChange() {
	configChan := c.config.WatchConfig()

	for {
		select {
		case newConfig := <-configChan:
			// 更新配置指针并重新连接
			c.mu.Lock()
			c.config = newConfig
			c.mu.Unlock()

			if err := c.initConnection(); err != nil {
				log.Printf("failed to reconnect after config change: %v", err)
			}
		case <-c.watchDone:
			// 收到退出信号，停止监听
			log.Println("cache config watcher stopped")
			return
		}
	}
}

// Close 优雅关闭缓存连接
// 实现 io.Closer 接口
func (c *Cache) Close() error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// 停止配置监听协程
	if c.watchDone != nil {
		close(c.watchDone)
		c.watchDone = nil
	}

	// 关闭 Redis 连接
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// GetClient 获取 Redis 客户端
// 线程安全：使用读锁保护
func (c *Cache) GetClient() *redis.Client {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.conn
}

// IsConnected 检查连接状态
// 线程安全：使用读锁保护
func (c *Cache) IsConnected() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.conn == nil {
		return false
	}

	return c.conn.Ping().Err() == nil
}

var ProviderSet = wire.NewSet(NewCache)
