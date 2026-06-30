package expired_strategy

import (
	"fmt"

	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/v2/config"
)

func NewExpiredStrategyImpl(config config.ExpiredConfig) IExpiredStrategy {
	var strategy IExpiredStrategy
	switch {
	case config.Name == "LRU":
		strategy = NewLRU(config.Len)
	case config.Name == "FIFO":
		strategy = NewFifo(config.Len)
	case config.Name == "FIFO_RING":
		strategy = NewFifoRingList(config.Len)
	default:
		panic(fmt.Sprintf("expire [%s] not support!", config.Name))
	}
	return strategy
}

var ProviderSet = wire.NewSet(NewExpiredStrategyImpl)
