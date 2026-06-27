package expired_strategy

import (
	"sync"
)

type entity struct {
	prev  *entity
	value string
	next  *entity
}

type LRU struct {
	mu sync.Mutex
	// 首节点
	node *entity
}

func NewLRU() *LRU {
	return &LRU{}
}

func (l *LRU) Push(v string) {

	l.mu.Lock()
	defer l.mu.Unlock()

	nn := &entity{
		value: v,
		// prev:  last,
		next: l.node,
	}

	if l.node == nil {
		l.node = nn
		l.node.next = nn
		l.node.prev = nn
		return
	}
	// 双向链表插入数据， 应该
	// 1. 首节点修改 prev 到最新节点
	// 2. 新节点修改 next 到首节点
	// 3. 之前的尾节点修改 next 到新节点

	// 拿到尾部
	last := l.node.prev
	nn.prev = last
	last.next = nn
	l.node.prev = nn
}

func (l *LRU) Pop(n int) []string {
	ret := make([]string, 0, n)
	if n <= 0 {
		return ret
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	if l.node == nil {
		return ret
	}
	for i := 0; i < n; i++ {

		// 移除最后一个
		// 1. 首节点的 prev 要换成倒数第二个
		// 2. 倒数第二个节点的 next 要换成首节点

		// 拿到最后一个
		last := l.node.prev
		if last == l.node {
			ret = append(ret, last.value)
			l.node = nil
			break
		}
		// 续上链表
		node := last.prev
		l.node.prev = node
		node.next = l.node

		ret = append(ret, last.value)
	}

	return ret
}

func (l *LRU) MoveToHead(entity *entity) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 需要干什么
	// 1. 改其 next 为当前首指针
	// 2. 将其 prev 为当前尾指针
	// 3. 其原始 prev 的 next 为其 next
	// 4. 其原始 next 的 prev 为其 prev

	// 无事发生
	if entity == l.node {
		return
	}

	c_prev := entity.prev
	c_next := entity.next

	c_prev.next = c_next
	c_next.prev = c_prev

	// 当前对尾部
	c_last := l.node.prev
	entity.prev = c_last
	entity.next = l.node
	c_last.next = entity
	l.node.prev = entity
	l.node = entity
}
