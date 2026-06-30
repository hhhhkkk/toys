package service

import (
	"github.com/google/wire"
	consistencyhash "github.com/hhhhkkk/mini-blog/v2/internal/service/consistency_hash"
	"github.com/hhhhkkk/mini-blog/v2/internal/service/expired_strategy"
)

var ProviderSet = wire.NewSet(
	expired_strategy.ProviderSet,
	consistencyhash.ProviderSet,
)
