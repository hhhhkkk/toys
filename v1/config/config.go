package config

import (
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type observer func(config Config)

type observerList struct {
	observers []observer
}

func (ol *observerList) AddObserver(fn observer) {
	ol.observers = append(ol.observers, fn)
}

func (ol *observerList) Notify(config Config) {
	for _, fn := range ol.observers {
		go fn(config)
	}
}

var (
	mu sync.Mutex
	ol = &observerList{}
)

func AddObserver(fn observer) {
	mu.Lock()
	defer mu.Unlock()
	ol.AddObserver(fn)
}

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
	Driver   string `json:"driver" yaml:"driver"`
	Dsn      string `json:"dsn" yaml:"dsn"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	Log    LogConfig    `yml:"log"`
	Cache  CacheConfig  `yaml:"cache"`
	Db     DB           `yaml:"db"`
}

type DB struct {
	Dsn string `yaml:"dsn"`
}

// 项目根目录
func GetRootPath() string {
	var root_path string

	if root_path == "" {
		root_path, _ = os.Getwd()
	}
	return root_path
}

func NewServerConfig() Config {
	cr := viper.New()

	cr.AddConfigPath("./")
	cr.AddConfigPath("./config")
	cr.AddConfigPath("./v1/config")
	cr.SetConfigName("app")
	cr.SetConfigType("yml")

	if err := cr.ReadInConfig(); err != nil {
		panic(err)
	}
	var config Config
	if err := cr.Unmarshal(&config); err != nil {
		panic(err)
	}

	go func() {
		cr.OnConfigChange(func(in fsnotify.Event) {
			if err := cr.Unmarshal(&config); err != nil {
				panic(err)
			}
			ol.Notify(config)
		})
		cr.WatchConfig()
	}()
	return config
}
