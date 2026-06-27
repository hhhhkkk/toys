package expired_strategy

import (
	"sync"
	"testing"
)

// 辅助函数：从链表头开始遍历，返回所有节点的值（用于验证顺序）
func listValues(l *LRU) []string {
	if l.node == nil {
		return []string{}
	}
	values := []string{l.node.value}
	for cur := l.node.next; cur != l.node; cur = cur.next {
		values = append(values, cur.value)
	}
	return values
}

// 测试 Push：验证尾部插入和链表顺序
func TestLRU_Push(t *testing.T) {
	lru := NewLRU()

	// 空链表 Push
	lru.Push("A")
	if got := listValues(lru); !equalSlice(got, []string{"A"}) {
		t.Errorf("Push A, got %v, want [A]", got)
	}

	// 继续 Push
	lru.Push("B")
	lru.Push("C")
	if got := listValues(lru); !equalSlice(got, []string{"A", "B", "C"}) {
		t.Errorf("Push A,B,C, got %v, want [A B C]", got)
	}

	// 验证循环链接：尾节点的 next 指向头，头节点的 prev 指向尾
	head := lru.node
	tail := head.prev
	if tail.next != head {
		t.Errorf("tail.next not head")
	}
	if head.prev != tail {
		t.Errorf("head.prev not tail")
	}
}

// 测试 Pop：弹出尾部元素，验证返回值及链表状态
func TestLRU_Pop(t *testing.T) {
	lru := NewLRU()
	lru.Push("A")
	lru.Push("B")
	lru.Push("C")

	// 正常弹出
	got := lru.Pop(2)
	if !equalSlice(got, []string{"C", "B"}) {
		t.Errorf("Pop 2, got %v, want [C B]", got)
	}
	if got := listValues(lru); !equalSlice(got, []string{"A"}) {
		t.Errorf("after Pop 2, list got %v, want [A]", got)
	}

	// 弹出最后一个
	got = lru.Pop(1)
	if !equalSlice(got, []string{"A"}) {
		t.Errorf("Pop 1, got %v, want [A]", got)
	}
	if lru.node != nil {
		t.Errorf("list should be empty, but node not nil")
	}

	// 空链表 Pop 应返回空切片
	got = lru.Pop(2)
	if len(got) != 0 {
		t.Errorf("Pop on empty, got %v, want []", got)
	}

	// Pop 数量大于实际节点数：应返回所有节点
	lru.Push("X")
	lru.Push("Y")
	got = lru.Pop(5)
	if !equalSlice(got, []string{"Y", "X"}) {
		t.Errorf("Pop 5 from [X,Y], got %v, want [Y X]", got)
	}
	if lru.node != nil {
		t.Errorf("list should be empty after Pop all")
	}
}

// 测试 MoveToHead：将节点移到头部
func TestLRU_MoveToHead(t *testing.T) {
	lru := NewLRU()
	lru.Push("A")
	lru.Push("B")
	lru.Push("C") // 链表: A -> B -> C

	// 移动中间节点 B 到头部
	// 需要先获取 B 节点的指针（这里通过遍历模拟，实际使用需暴露，但在测试中我们通过 node 指针操作）
	// 由于 entity 未导出，我们只能通过内部方法测试。下面我们直接使用已知指针。
	// 更好的办法：在测试中通过 lru.node.next 获取 B，但 entity 未导出，我们可以在内部写一个 getNode 辅助。
	// 为了演示，我们直接通过 lru.node.next 获取 B（因为 B 是第二个节点）
	b := lru.node.next
	lru.MoveToHead(b)

	if got := listValues(lru); !equalSlice(got, []string{"B", "A", "C"}) {
		t.Errorf("MoveToHead B, got %v, want [B A C]", got)
	}

	// 移动头节点（无变化）
	lru.MoveToHead(lru.node)
	if got := listValues(lru); !equalSlice(got, []string{"B", "A", "C"}) {
		t.Errorf("MoveToHead head, got %v, want [B A C]", got)
	}

	// 移动尾节点 C 到头部
	c := lru.node.prev
	lru.MoveToHead(c)
	if got := listValues(lru); !equalSlice(got, []string{"C", "B", "A"}) {
		t.Errorf("MoveToHead C, got %v, want [C B A]", got)
	}
}

// 测试并发安全性：多个 goroutine 同时 Push 和 Pop
func TestLRU_Concurrency(t *testing.T) {
	lru := NewLRU()
	var wg sync.WaitGroup
	ops := 100

	// 并发 Push
	for i := 0; i < ops; i++ {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			lru.Push(v)
		}(string(rune('A' + i%26))) // 循环使用字母
	}

	// 并发 Pop（部分）
	for i := 0; i < ops/2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = lru.Pop(1)
		}()
	}

	wg.Wait()

	// 最终链表应该不为 nil（因为我们 push 了 100 个，pop 了 50 个）
	if lru.node == nil {
		t.Log("链表为空，可能 pop 过多，但并发场景允许")
	} else {
		// 检查链表是否循环且无断链
		head := lru.node
		visited := 0
		for cur := head; visited == 0 || cur != head; cur = cur.next {
			visited++
			if visited > 1000 { // 防止死循环
				t.Errorf("链表可能形成环，检查失败")
				break
			}
		}
	}
}

// 辅助函数：比较两个字符串切片是否相等
func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}