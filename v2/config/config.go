package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ConsistencyConfig struct {
	Name string `mapstructure:"name"`
	Len  int    `mapstructure:"len"`
}

type ExpiredConfig struct {
	Name string `mapstructure:"name"`
	Len  int    `mapstructure:"len"`
}

type Host struct {
	Name     string `yml:"name"`
	FakerNum int    `yml:"fakerNum"`
	Host     string `yml:"host"`
}

type Config struct {
	AppName           string            `mapstructure:"appName"`
	ExpiredConfig     ExpiredConfig     `mapstructure:"expireConfig"`
	ConsistencyConfig ConsistencyConfig `mapstructure:"consistencyConfig"`
	HostList          []Host            `mapstructure:"replicas"`
}

func New() Config {
	confReader := viper.New()
	confReader.AddConfigPath("./../v2/config")
	confReader.AddConfigPath("./../v2/config")
	confReader.AddConfigPath("./../config")
	confReader.AddConfigPath("./config")
	confReader.AddConfigPath("./v2/config")
	confReader.SetConfigType("yml")
	confReader.SetConfigName("config")

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("当前工作目录:", dir)

	if err := confReader.ReadInConfig(); err != nil {
		panic(err)
	}
	var config Config
	confReader.Unmarshal(&config)
	return config
}

func NewExpiredConfig(config Config) ExpiredConfig {
	return config.ExpiredConfig
}

func NewConsistencyConfig(config Config) ConsistencyConfig {
	return config.ConsistencyConfig
}

func NewHostList(config Config) []Host {
	return config.HostList
}
