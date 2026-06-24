package data

import (
	"log"
	"sync/atomic"

	"github.com/go-redis/redis/v7"
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/v1/config"
	"github.com/hhhhkkk/mini-blog/v1/internal/biz/repository/admin"
	"github.com/hhhhkkk/mini-blog/v1/internal/biz/repository/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Cache Redis 缓存客户端封装
// 提供配置热重载和优雅关闭支持
type Cache struct {
	conn atomic.Value
}

// NewCache 创建缓存实例并初始化连接
// 启动配置变更监听协程
func NewCache(cfg config.Config) *Cache {

	client, err := initConnection(cfg)

	if err != nil {
		panic(err)
	}
	c := &Cache{}
	c.conn.Store(client)
	config.AddObserver(func(newConfig config.Config) {
		// 创建新连接
		newClient, err := initConnection(newConfig)
		if err != nil {
			log.Printf("failed to connect to redis: %v", err)
			return
		}
		log.Printf("redis reconnected successfully: %s", newClient.Options().Addr)

		oldValue := c.conn.Swap(newClient)
		// 关闭旧连接
		if ov, ok := oldValue.(*redis.Client); ok {
			ov.Close()
		}
	})
	return c
}

// initConnection 初始化 Redis 连接
func initConnection(cfg config.Config) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.Dsn,
		DB:       cfg.Cache.DB,
		Password: cfg.Cache.Password,
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
	return c.conn.Load().(*redis.Client)
}

type DB struct {
	conn atomic.Value
}

func (db *DB) GetClient() *gorm.DB {
	return db.conn.Load().(*gorm.DB)
}

func NewDB(cfg config.Config) *DB {
	ret := &DB{}
	db := initDB(cfg)
	ret.conn.Store(db)
	config.AddObserver(func(newCfg config.Config) {
		old_db := ret.conn.Swap(initDB(newCfg))
		if old, ok := old_db.(*gorm.DB); ok {
			db, _ := old.DB()
			db.Close()
		}
	})
	return ret
}

func initDB(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: cfg.Db.Dsn,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

var ProviderSet = wire.NewSet(
	wire.Bind(new(admin.Repo), new(*AdminUserRepoImpl)),
	wire.Bind(new(api.UserRepo), new(*UserRepoImpl)),
	wire.Bind(new(api.UserIdentityRepo), new(*IdentityRepoImpl)),
	wire.Bind(new(api.InviteRepo), new(*InviteRecordRepoImpl)),
	wire.Bind(new(api.InviteAwardRepo), new(*InviteAwardRepoImpl)),
	NewCache,
	NewDB,
	NewAdminUserRepoImpl,
	NewUserRepoImpl,
	NewIdentityRepoImpl,
	NewInviteRecordRepoImpl,
	NewInviteAwardRepoImpl,
)
