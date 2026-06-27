package cache

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/v2/internal/service/expired_strategy"
)

func ProvideLRU() expired_strategy.IExpiredStrategy {
	return expired_strategy.NewLRU(3)
}

var ProviderSet = wire.NewSet(
	NewCacheService,
	ProvideLRU,
)
