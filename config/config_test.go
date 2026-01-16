package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hhhhkkk/mini-blog/config"
)

func TestNewAppConfig(t *testing.T) {
	appConfig := config.NewAppConfig()

	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"app name", appConfig.Name, ""},
		{"app version", appConfig.Version, ""},
		{"app port", appConfig.Port, 0},
		{"app host", appConfig.Host, ""},
		{"app env", appConfig.Env, ""},
	}

	for _, tt := range tests {
		if tt.got == tt.expected {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, tt.got)
		}
	}
}

func TestGetRootPath(t *testing.T) {
	path := config.GetRootPath()

	// 验证路径不为空
	if path == "" {
		t.Error("root path should not be empty")
	}

	// 验证路径是绝对路径
	if !filepath.IsAbs(path) {
		t.Errorf("root path should be absolute path, got: %s", path)
	}

	// 验证路径存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("root path should exist, got: %s", path)
	}
}
