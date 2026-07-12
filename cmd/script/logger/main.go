package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Server ServerConfig `yaml:"server"`
	Log    LogConfig    `yaml:"log"`
	Cache  CacheConfig  `yaml:"cache"`
}

type ServerConfig struct {
	Name    string `yaml:"name" mapstructure:"name"`
	Version string `yaml:"version" mapstructure:"version"`
	Host    string `yml:"host"`
	Port    string `yml:"port"`
	Env     string `yml:"port"`
}

type LogConfig struct {
	ErrFile string `yaml:"err_file" mapstructure:"err_file"`
}

type CacheConfig struct {
	Dsn string `yaml:"dsn"`
	Pwd string `yaml:"password"`
	Db  int    `json:"db" yaml:"db"`
}

type observer struct {
	lock      sync.RWMutex
	observers []func(AppConfig)
}

func (ob *observer) AddObserver(fn func(AppConfig)) {
	ob.lock.Lock()
	defer ob.lock.Unlock()
	ob.observers = append(ob.observers, fn)
}

func AddObserver(fn func(AppConfig)) {
	fmt.Println("add observer")
	ob.AddObserver(fn)
}

var (
	config = AppConfig{}
	ch     = make(chan AppConfig, 1)
	ob     = &observer{
		lock: sync.RWMutex{},
	}
)

func init() {

	customerViper := viper.New()
	customerViper.AddConfigPath("./cmd/script/logger/")
	customerViper.SetConfigName("config")
	customerViper.SetConfigType("yaml")

	if err := customerViper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := customerViper.Unmarshal(&config); err != nil {
		panic(err)
	}

	customerViper.OnConfigChange(func(in fsnotify.Event) {
		if err := customerViper.Unmarshal(&config); err != nil {
			fmt.Println("config change error log" + err.Error())
			return
		}
		ch <- config
	})
	go func() {
		customerViper.WatchConfig()
	}()

	go func() {
		for {
			config := <-ch

			ob.lock.Lock()
			fmt.Println("config change:", config)
			var ss []func(AppConfig)
			ss = append(ss, ob.observers...)
			copy(ss, ob.observers)
			ob.lock.Unlock()

			// 在锁外执行回调
			for _, fn := range ss {
				go fn(config)
			}
		}
	}()
}

func GetAppConfig() AppConfig {
	return config
}

func GetServerConfig() ServerConfig {
	return config.Server
}

func GetLogConfig() LogConfig {
	return config.Log
}

func GetCacheConfig() CacheConfig {
	return config.Cache
}

func main() {
	config := GetAppConfig()
	fmt.Println(config)

	go func() {

		fmt.Println("start work...")
		cfg := GetCacheConfig()
		redisClient := redis.NewClient(&redis.Options{
			Addr: cfg.Dsn,
			DB:   cfg.Db,
		})

		AddObserver(func(ac AppConfig) {
			redisClient.Close()

			redisClient = redis.NewClient(&redis.Options{
				Addr: ac.Cache.Dsn,
				DB:   ac.Cache.Db,
			})
			fmt.Println(ac.Cache.Db)
		})
		fmt.Println("new tick")
		tick := time.NewTicker(3 * time.Second)

		for {
			<-tick.C
			fmt.Println("tick...")
			name, err := redisClient.Get("user1").Result()
			if err != nil {
				fmt.Println("get error:", err)
				continue
			}
			fmt.Println("get name:", name)
		}
	}()

	select {}
}
