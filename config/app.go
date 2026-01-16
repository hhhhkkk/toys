package config

import (
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name      string `json:"name" yaml:"name"`
	Version   string `json:"version" yaml:"version"`
	Port      int    `json:"port" yaml:"port"`
	Host      string `json:"host" yaml:"host"`
	ErrorPath string `json:"error_path" yaml:"error_path"`
	Env       string `json:"env" yaml:"env"`
}

var root_path string

// 项目根目录
func GetRootPath() string {
	if root_path == "" {
		root_path, _ = os.Getwd()
	}
	return root_path
}

func NewAppConfig() *AppConfig {
	cr := viper.New()

	cr.AddConfigPath("./")
	cr.AddConfigPath("./config")
	cr.SetConfigName("app")
	cr.SetConfigType("yml")

	if err := cr.ReadInConfig(); err != nil {
		panic(err)
	}
	var appConfig AppConfig
	if err := cr.Unmarshal(&appConfig); err != nil {
		panic(err)
	}
	return &appConfig
}
