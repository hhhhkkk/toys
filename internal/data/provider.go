package data

import (
	"log"

	"github.com/go-redis/redis/v7"
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/config"
)

// Cache Redis 缓存客户端封装
// 提供配置热重载和优雅关闭支持
type Cache struct {
	conn *redis.Client
}

// NewCache 创建缓存实例并初始化连接
// 启动配置变更监听协程
func NewCache(cfg config.Config) *Cache {

	client, err := initConnection(cfg)

	if err != nil {
		panic(err)
	}
	c := &Cache{
		conn: client,
	}
	config.AddObserver(func(config config.Config) {
		client, err := initConnection(config)
		if err != nil {
			log.Printf("failed to connect to redis: %v", err)
			return
		}
		log.Printf("redis connected successfully: %s", client.Options().Addr)

		c.conn = client
	})
	return c
}

// initConnection 初始化 Redis 连接
func initConnection(cfg config.Config) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr: cfg.Cache.Dsn,
		DB:   cfg.Cache.DB,
	})

	// 验证连接
	if err := conn.Ping().Err(); err != nil {
		log.Printf("failed to connect to redis: %v", err)
		return nil, err
	}
	log.Printf("redis connected successfully: %s", conn.Options().Addr)
	return conn, nil
}

func (c *Cache) GetClient() *redis.Client {
	return c.conn
}

var ProviderSet = wire.NewSet(NewCache)
