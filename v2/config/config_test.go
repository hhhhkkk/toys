package config

import (
	"testing"
)

func TestGetAppName(t *testing.T) {
	conf := New()
	if conf.AppName != "v2" {
		t.Errorf("app name test fail, got `%s`", conf.AppName)
	}
}

func TestGetSubConfig(t *testing.T) {
	conf := New()

	var zero ConsistencyConfig
	if conf.ConsistencyConfig == zero {
		t.Error("get ConsistencyConfig is nil")
	}

	var zeroE ExpiredConfig
	if conf.ExpiredConfig == zeroE {
		t.Error("get ExpiredConfig is nil")
	}

	if conf.HostList == nil {
		t.Error("host list is empty!")
	}
}
