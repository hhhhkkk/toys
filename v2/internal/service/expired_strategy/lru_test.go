package expired_strategy

import (
	"sync"
	"testing"
)

// 辅助函数：从链表头开始遍历，返回所有节点的值（用于验证链表顺序）
func listValues(l *LRU) []string {
	if l.node == nil {
		return []string{}
	}
	vals := []string{l.node.value}
	for cur := l.node.next; cur != l.node; cur = cur.next {
		vals = append(vals, cur.value)
	}
	return vals
}

// 辅助：比较两个字符串切片是否相等（忽略顺序仅用于长度或内容）
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

func TestLRU_EMPTY_POP(t *testing.T) {
	lru := NewLRU(3)
	for i := 0; i < 100; i++ {
		_, ok := lru.Pop()
		if ok {
			t.Error("empty lru pop ok")
		}
	}
}

// 测试 Push 基本插入（无淘汰）
func TestLRU_Push_Basic(t *testing.T) {
	lru := NewLRU(3)
	evicted, ok := lru.Push("A")
	if ok || evicted != "" {
		t.Errorf("Push A expected no eviction, got evicted=%q, ok=%v", evicted, ok)
	}
	// 链表应为 [A]
	if got := listValues(lru); !equalSlice(got, []string{"A"}) {
		t.Errorf("list = %v, want [A]", got)
	}

	lru.Push("B")
	lru.Push("C")
	// 链表应为 [C, B, A]（新插入在头部）
	if got := listValues(lru); !equalSlice(got, []string{"C", "B", "A"}) {
		t.Errorf("list = %v, want [C, B, A]", got)
	}

	if lru.Len() != 3 {
		t.Errorf("Len = %d, want 3", lru.Len())
	}
}

// 测试 Push 触发容量淘汰（淘汰尾部最久未使用）
func TestLRU_Push_Eviction(t *testing.T) {
	lru := NewLRU(2)
	lru.Push("A") // [A]
	lru.Push("B") // [B, A]

	// 插入 C，容量 2，应淘汰尾部 A
	evicted, ok := lru.Push("C")
	if !ok || evicted != "A" {
		t.Errorf("Push C expected evict A, got evicted=%q, ok=%v", evicted, ok)
	}
	if got := listValues(lru); !equalSlice(got, []string{"C", "B"}) {
		t.Errorf("after eviction list = %v, want [C, B]", got)
	}
	if lru.Len() != 2 {
		t.Errorf("Len = %d, want 2", lru.Len())
	}

	// 继续插入 D，淘汰 B
	evicted, ok = lru.Push("D")
	if !ok || evicted != "B" {
		t.Errorf("Push D expected evict B, got evicted=%q, ok=%v", evicted, ok)
	}
	if got := listValues(lru); !equalSlice(got, []string{"D", "C"}) {
		t.Errorf("after eviction list = %v, want [D, C]", got)
	}
}

// 测试 Push 容量为 1 的极端情况
func TestLRU_Push_CapacityOne(t *testing.T) {
	lru := NewLRU(1)
	evicted, ok := lru.Push("A")
	if ok {
		t.Errorf("expected no eviction, got evicted")
	}
	// 再次 Push 应淘汰 A
	evicted, ok = lru.Push("B")
	if !ok || evicted != "A" {
		t.Errorf("expected evict A, got %q", evicted)
	}
	if got := listValues(lru); !equalSlice(got, []string{"B"}) {
		t.Errorf("list = %v, want [B]", got)
	}
	if lru.Len() != 1 {
		t.Errorf("Len = %d, want 1", lru.Len())
	}
}

// 测试 Pop 基本功能（弹出尾部）
func TestLRU_Pop_Basic(t *testing.T) {
	lru := NewLRU(5)
	lru.Push("A")
	lru.Push("B")
	lru.Push("C") // [C, B, A]

	val, ok := lru.Pop()
	if !ok || val != "A" {
		t.Errorf("Pop got %q, %v, want 'A', true", val, ok)
	}
	if got := listValues(lru); !equalSlice(got, []string{"C", "B"}) {
		t.Errorf("after Pop list = %v, want [C, B]", got)
	}
	if lru.Len() != 2 {
		t.Errorf("Len = %d, want 2", lru.Len())
	}

	// 继续 Pop 直到空
	lru.Pop()
	lru.Pop()
	if lru.Len() != 0 {
		t.Errorf("Len should be 0, got %d", lru.Len())
	}
	if lru.node != nil {
		t.Errorf("node should be nil")
	}

	// 空链表 Pop
	val, ok = lru.Pop()
	if ok || val != "" {
		t.Errorf("Pop on empty got %q, %v, want '', false", val, ok)
	}
}

// 测试 Pop 与 Push 混合操作，验证 LRU 顺序正确
func TestLRU_PushPop_Mixed(t *testing.T) {
	lru := NewLRU(3)
	lru.Push("A")
	lru.Push("B")
	lru.Push("C") // [C, B, A]

	// 手动 Pop 掉 A（尾部）
	lru.Pop() // 弹出 A
	// 当前 [C, B]
	if got := listValues(lru); !equalSlice(got, []string{"C", "B"}) {
		t.Errorf("after pop list = %v, want [C, B]", got)
	}

	// 再 Push D，容量 3，当前长度 2，不会淘汰
	evicted, ok := lru.Push("D")
	if ok {
		t.Errorf("expected no eviction, got evicted %q", evicted)
	}
	// 链表应为 [D, C, B]
	if got := listValues(lru); !equalSlice(got, []string{"D", "C", "B"}) {
		t.Errorf("after Push D list = %v, want [D, C, B]", got)
	}

	// 现在长度=3，Push E 会触发淘汰，淘汰尾部 B
	evicted, ok = lru.Push("E")
	if !ok || evicted != "B" {
		t.Errorf("Push E expected evict B, got %q, ok=%v", evicted, ok)
	}
	// 链表应为 [E, D, C]（因为淘汰了尾部的 B）
	if got := listValues(lru); !equalSlice(got, []string{"E", "D", "C"}) {
		t.Errorf("after Push E list = %v, want [E, D, C]", got)
	}

	// 再 Push F，当前长度=3，又触发淘汰，淘汰尾部 C
	evicted, ok = lru.Push("F")
	if !ok || evicted != "C" {
		t.Errorf("Push F expected evict C, got %q, ok=%v", evicted, ok)
	}
	// 最终链表应为 [F, E, D]
	if got := listValues(lru); !equalSlice(got, []string{"F", "E", "D"}) {
		t.Errorf("final list = %v, want [F, E, D]", got)
	}
}

// 测试并发安全：多个 goroutine 同时 Push 和 Pop
func TestLRU_Concurrency(t *testing.T) {
	lru := NewLRU(10)
	var wg sync.WaitGroup
	const goroutines = 10
	const opsPerGoroutine = 100

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			base := string(rune('A' + id))
			for j := 0; j < opsPerGoroutine; j++ {
				// 交替 Push 和 Pop
				if j%2 == 0 {
					lru.Push(base + string(rune('a'+j%26)))
				} else {
					lru.Pop()
				}
			}
		}(i)
	}
	wg.Wait()

	// 最终链表可能非空，但应保持循环结构且无 nil 指针
	if lru.node != nil {
		// 遍历检查循环完整性
		head := lru.node
		visited := 0
		cur := head
		for {
			visited++
			if visited > 10000 { // 防无限循环
				t.Errorf("possible cycle detected, visited > 10000")
				break
			}
			if cur.next == nil || cur.prev == nil {
				t.Errorf("nil pointer in linked list")
				break
			}
			cur = cur.next
			if cur == head {
				break
			}
		}
		// Len 应不大于 maxLen
		if lru.Len() > lru.maxLen {
			t.Errorf("Len = %d > maxLen %d", lru.Len(), lru.maxLen)
		}
	}
}

// 测试 Len 方法在并发环境下的正确性（仅读取）
func TestLRU_LenConcurrency(t *testing.T) {
	lru := NewLRU(5)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lru.Push("x")
			_ = lru.Len()
		}()
	}
	wg.Wait()
	// 最终长度应不超过5
	if lru.Len() > 5 {
		t.Errorf("Len = %d > 5", lru.Len())
	}
}

// 测试 NewLRU 传入非法容量（<=0）应 panic
func TestNewLRU_InvalidCapacity(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for maxLen<=0")
		}
	}()
	NewLRU(0)
}
