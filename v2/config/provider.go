package config

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewExpiredConfig,
	NewConsistencyConfig,
	NewHostList,
	New,
)
