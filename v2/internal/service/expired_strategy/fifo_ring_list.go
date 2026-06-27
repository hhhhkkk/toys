package expired_strategy

import (
	"sync"
)

// FifoRingList 利用固定长度数组， 只移动下标，减少 append 分片的操作
type FifoRingList struct {
	maxLen int
	data   []string
	lock   sync.Mutex
	head   int
	tail   int
	len    int
}

func NewFifoRingList(maxLen int) *FifoRingList {
	if maxLen <= 0 {
		panic("FIFO not support max_len = 0.")
	}
	return &FifoRingList{
		maxLen: maxLen,
		data:   make([]string, maxLen),
	}
}

func (f *FifoRingList) Len() int {
	return f.len
}

func (f *FifoRingList) Push(k string) (string, bool) {
	f.lock.Lock()
	defer f.lock.Unlock()

	ret := ""
	pop := false
	if f.len == f.maxLen {
		pop = true
		ret = f.data[f.head]
		f.head = (f.head + 1) % f.maxLen
		f.len--
	}

	f.data[f.tail] = k
	f.tail = (f.tail + 1) % f.maxLen

	f.len++
	return ret, pop
}

func (f *FifoRingList) Pop() (string, bool) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.len == 0 {
		return "", false
	}
	ret := f.data[f.head]
	// 帮助 GC 回收字符串引用（非必须）
	f.data[f.head] = ""
	f.head = (f.head + 1) % f.maxLen
	f.len--
	return ret, true
}
