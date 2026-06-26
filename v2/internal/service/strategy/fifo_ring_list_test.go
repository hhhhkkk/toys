package strategy

import (
	"sync"
	"testing"
)

// TestFifoRingList_New 测试构造函数
func TestFifoRingList_New(t *testing.T) {
	t.Run("valid maxLen", func(t *testing.T) {
		f := NewFifoRingList(3)
		if f.maxLen != 3 {
			t.Errorf("expected maxLen=3, got %d", f.maxLen)
		}
		if f.len != 0 {
			t.Errorf("expected empty data, got len=%d", len(f.data))
		}
	})

	t.Run("maxLen <= 0 should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for maxLen <= 0")
			}
		}()
		NewFifoRingList(0)
	})
}

// TestFifoRingList_Push 测试入队逻辑
func TestFifoRingList_Push(t *testing.T) {
	tests := []struct {
		name          string
		maxLen        int
		pushes        []string
		expectedEvict []string // 每次 Push 返回的淘汰值，nil 表示无需校验
		expectedData  []string // 最终队列内容
	}{
		{
			name:          "simple push without eviction",
			maxLen:        3,
			pushes:        []string{"a", "b", "c"},
			expectedEvict: []string{"", "", ""},
			expectedData:  []string{"a", "b", "c"},
		},
		{
			name:          "push with eviction",
			maxLen:        3,
			pushes:        []string{"a", "b", "c", "d", "e"},
			expectedEvict: []string{"", "", "", "a", "b"},
			expectedData:  []string{"c", "d", "e"},
		},
		{
			name:          "push empty string",
			maxLen:        2,
			pushes:        []string{"", "x", ""},
			expectedEvict: []string{"", "", ""},
			expectedData:  []string{"x", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFifoRingList(tt.maxLen)
			for i, val := range tt.pushes {
				evicted, _ := f.Push(val)
				if tt.expectedEvict[i] != evicted {
					t.Errorf("Push %d: expected evicted=%q, got %q", i, tt.expectedEvict[i], evicted)
				}
			}

			// 检查最终队列内容
			if len(f.data) != len(tt.expectedData) {
				t.Errorf("final length: expected %d, got %d", len(tt.expectedData), len(f.data))
			}
			for i, expected := range tt.expectedData {
				val, ok := f.Pop()
				if !ok {
					t.Errorf("Pop %d: expected %q, but queue is empty", i, expected)
				} else if val != expected {
					t.Errorf("Pop %d: expected %q, got %q", i, expected, val)
				}
			}
		})
	}
}

// TestFifoRingList_Pop 测试出队逻辑
func TestFifoRingList_Pop(t *testing.T) {
	f := NewFifoRingList(3)
	f.Push("a")
	f.Push("b")
	f.Push("c")

	// 正常出队
	val, ok := f.Pop()
	if !ok || val != "a" {
		t.Errorf("Pop expected a, got %q (ok=%v)", val, ok)
	}
	val, ok = f.Pop()
	if !ok || val != "b" {
		t.Errorf("Pop expected b, got %q (ok=%v)", val, ok)
	}

	// 再 Push 一个，验证队列继续工作
	f.Push("d")
	val, ok = f.Pop()
	if !ok || val != "c" {
		t.Errorf("Pop expected c, got %q (ok=%v)", val, ok)
	}
	val, ok = f.Pop()
	if !ok || val != "d" {
		t.Errorf("Pop expected d, got %q (ok=%v)", val, ok)
	}

	// 空队列 Pop
	val, ok = f.Pop()
	if ok || val != "" {
		t.Errorf("Pop from empty queue: expected \"\", false, got %q, %v", val, ok)
	}
}

// TestFifoRingList_Concurrency 并发测试（检测数据竞争）
func TestFifoRingList_Concurrency(t *testing.T) {
	f := NewFifoRingList(100)
	const goroutines = 50
	const iterations = 1000
	var wg sync.WaitGroup
	wg.Add(goroutines * 2)

	// 并发 Push
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				f.Push("x")
			}
		}(i)
	}

	// 并发 Pop
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				f.Pop()
			}
		}(i)
	}

	wg.Wait()
	// 如果没死锁或 panic，测试通过
}

// BenchmarkFifo_Push 测试入队性能
func BenchmarkFifoRingList_Push(b *testing.B) {
	f := NewFifoRingList(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Push("test")
	}
}

// BenchmarkFifo_PushWithEviction 测试满队列时入队性能
func BenchmarkFifoRingList_PushWithEviction(b *testing.B) {
	f := NewFifoRingList(100)
	// 先填满
	for i := 0; i < 100; i++ {
		f.Push("init")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Push("test")
	}
}

// BenchmarkFifo_Pop 测试出队性能
func BenchmarkFifoRingList_Pop(b *testing.B) {
	f := NewFifoRingList(1000)
	// 先填满
	for i := 0; i < 1000; i++ {
		f.Push("init")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Pop()
		// 保持队列不空，避免 Pop 返回 false 影响性能
		if i%1000 == 0 {
			f.Push("refill")
		}
	}
}

// BenchmarkFifo_Mixed 混合 Push/Pop（模拟真实场景）
func BenchmarkFifoRingList_Mixed(b *testing.B) {
	f := NewFifoRingList(100)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			f.Push("test")
			f.Pop()
		}
	})
}
