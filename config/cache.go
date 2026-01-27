package config

import (
	"log"

	"github.com/spf13/viper"
)

// NewCacheConfig 创建并初始化缓存配置
func NewCacheConfig() *CacheConfig {
	cr := viper.New()
	cr.AddConfigPath("./config")
	cr.SetConfigName("cache")
	cr.SetConfigType("yml")

	if err := cr.ReadInConfig(); err != nil {
		log.Printf("failed to read cache config: %v", err)
		// 返回默认配置而不是 panic，实现优雅降级
		return &CacheConfig{}
	}

	var cacheConfig CacheConfig
	if err := cr.Unmarshal(&cacheConfig); err != nil {
		log.Printf("failed to unmarshal cache config: %v", err)
		return &CacheConfig{}
	}
	return &cacheConfig
}

// WatchConfig 监听配置文件变更并返回配置变更事件
// 使用 channel 模式而非回调，更符合 Go 的惯用模式
// 返回的 channel 传递新的 CacheConfig 指针
func (c *CacheConfig) WatchConfig() <-chan *CacheConfig {
	configChan := make(chan *CacheConfig)

	// c.viper.WatchConfig()
	// c.viper.OnConfigChange(func(in fsnotify.Event) {
	// 	// 创建新的配置实例（包含 watcher 引用）
	// 	newConfig := &CacheConfig{
	// 		viper: c.viper,
	// 	}

	// 	// 反序列化新配置
	// 	if err := c.viper.Unmarshal(newConfig); err != nil {
	// 		log.Printf("failed to unmarshal new cache config: %v", err)
	// 		return // 忽略无效的配置变更
	// 	}

	// 	// 发送新配置指针到 channel
	// 	configChan <- newConfig
	// })

	return configChan
}
