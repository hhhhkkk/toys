package expired_strategy

import (
	"sync"

	"github.com/hhhhkkk/mini-blog/v2/pkg"
)

type Fifo struct {
	maxLen uint64
	data   []string
	lock   sync.Mutex
}

var _ IExpiredStrategy = (*Fifo)(nil)

func NewFifo(l int) *Fifo {
	if l <= 0 {
		panic("FIFO not support max_len = 0.")
	}
	return &Fifo{
		maxLen: uint64(l),
		data:   make([]string, l),
	}
}

func (s *Fifo) Push(k string) (string, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	ret := ""
	exists := false
	if s.maxLen == uint64(len(s.data)) {
		// 移除队首
		ret = s.data[0]
		s.data = s.data[1:]
		exists = true
	}

	s.data = append(s.data, k)
	return ret, exists
}

func (s *Fifo) Pop() (string, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.data) == 0 {
		return "", false
	}
	ret := s.data[0]
	s.data = s.data[1:]
	return ret, true
}

func (s *Fifo) Len() int {
	return pkg.ReduceSlice(s.data, 0, func(dv int, s string) int {
		incr := pkg.If(s == "", 0, 1)
		dv += incr
		return dv
	})
}
