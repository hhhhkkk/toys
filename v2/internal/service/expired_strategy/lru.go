package expired_strategy

import (
	"sync"
	"sync/atomic"
)

type entity struct {
	prev  *entity
	value string
	next  *entity
}

var _ IExpiredStrategy = (*LRU)(nil)

type LRU struct {
	mu sync.Mutex
	// 首节点
	node *entity

	maxLen int

	currentLen atomic.Int64
}

func NewLRU(maxLen int) *LRU {
	if maxLen <= 0 {
		panic("LRU not support max_len = 0.")
	}
	return &LRU{
		maxLen: maxLen,
	}
}

func (l *LRU) Len() int {
	return int(l.currentLen.Load())
}

func (l *LRU) Push(v string) (string, bool) {

	l.mu.Lock()
	defer l.mu.Unlock()

	nn := &entity{
		value: v,
		next:  l.node,
	}

	l.currentLen.Add(1)
	if l.node == nil {
		l.node = nn
		l.node.next = nn
		l.node.prev = nn
		return "", false
	} else {
		c_lasted := l.node.prev
		c_lasted.next = nn
		nn.next = l.node
		nn.prev = c_lasted
		l.node.prev = nn
		l.node = nn
	}

	ret := ""
	if int(l.currentLen.Load()) > l.maxLen {
		last := l.node.prev
		// 续上链表
		node := last.prev
		l.node.prev = node
		node.next = l.node

		ret = last.value
		l.currentLen.Add(-1)
		return ret, true
	}
	return ret, false
}

func (l *LRU) Pop() (string, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.node == nil {
		return "", false
	}

	ret := ""

	defer func(l *LRU) {
		l.currentLen.Add(-1)
	}(l)

	// 移除最后一个
	// 1. 首节点的 prev 要换成倒数第二个
	// 2. 倒数第二个节点的 next 要换成首节点

	// 拿到最后一个
	last := l.node.prev
	if last == l.node {
		ret = last.value
		l.node = nil
		return ret, true
	}
	// 续上链表
	node := last.prev
	l.node.prev = node
	node.next = l.node

	ret = last.value

	return ret, true
}
