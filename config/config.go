package config

import (
	"os"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name      string `json:"name" yaml:"name"`
	Version   string `json:"version" yaml:"version"`
	Port      int    `json:"port" yaml:"port"`
	Host      string `json:"host" yaml:"host"`
	ErrorPath string `json:"error_path" yaml:"error_path"`
	Env       string `json:"env" yaml:"env"`
	RootDir   string `json:"root_dir" yaml:"root_dir"`
}

type LogConfig struct {
	ErrLogPath string `json:"path" yaml:"path"`
}

// cacheItem 单个缓存实例配置
type CacheConfig struct {
	Driver string `json:"driver" yaml:"driver"`
	Host   string `json:"host" yaml:"host"`
	Port   int    `json:"port" yaml:"port"`
	Auth   string `json:"auth" yaml:"auth"`
	DB     int    `json:"db" yaml:"db"`
}

type Config struct {
	server ServerConfig `yaml:"server"`
	log    LogConfig    `yml:"log"`
	cache  CacheConfig  `yaml:"cache"`
}

// 项目根目录
func GetRootPath() string {
	var root_path string

	if root_path == "" {
		root_path, _ = os.Getwd()
	}
	return root_path
}

func NewServerConfig() *Config {
	cr := viper.New()

	cr.AddConfigPath("./")
	cr.AddConfigPath("./config")
	cr.SetConfigName("app")
	cr.SetConfigType("yml")

	if err := cr.ReadInConfig(); err != nil {
		panic(err)
	}
	var config Config
	if err := cr.Unmarshal(&config); err != nil {
		panic(err)
	}
	return &config
}
