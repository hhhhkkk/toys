package cache

import (
	"sync"
	"testing"

	consistencyhash "github.com/hhhhkkk/mini-blog/v2/internal/service/consistency_hash"
	"github.com/hhhhkkk/mini-blog/v2/internal/service/expired_strategy"
)

func defaultService() *CacheService {
	defaultExpiredStrategy := expired_strategy.NewFifo(8)
	hash := consistencyhash.New()
	return NewCacheService(defaultExpiredStrategy, hash)
}

// TestAdd 测试 add 函数
func TestAdd(t *testing.T) {
	cs := defaultService()
	key := "testKey"
	node := &Node{Value: "testValue", Size: 9}

	add(cs, key, node)

	// 验证是否成功添加
	if got, ok := cs.data[key]; !ok || got != node {
		t.Errorf("add failed: got %v, want %v", got, node)
	}

	// 覆盖已有键
	newNode := &Node{Value: "newValue", Size: 8}
	add(cs, key, newNode)
	if got := cs.data[key]; got != newNode {
		t.Errorf("add overwrite failed: got %v, want %v", got, newNode)
	}
}

// TestRemove 测试 remove 函数
func TestRemove(t *testing.T) {
	cs := defaultService()
	key := "testKey"
	node := &Node{Value: "testValue", Size: 9}

	// 先添加
	add(cs, key, node)

	// 删除存在的键
	remove(cs, key)
	if _, ok := cs.data[key]; ok {
		t.Errorf("remove failed: key still exists")
	}

	// 删除不存在的键（应安全，不 panic）
	remove(cs, "non-existing") // 不应该 panic
}

// TestAddRemoveConcurrency 并发测试，检测数据竞争
// 使用 go test -race 运行以验证锁的正确性
func TestAddRemoveConcurrency(t *testing.T) {

	cs := defaultService()
	const goroutines = 50
	const iterations = 100
	var wg sync.WaitGroup
	wg.Add(goroutines * 3) // 一半写，一半删

	// 写入 goroutines
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			key := "shared"
			node := &Node{Value: "value", Size: 5}
			for j := 0; j < iterations; j++ {
				add(cs, key, node)
			}
		}(i)
	}

	// 删除 goroutines
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			key := "shared"
			for j := 0; j < iterations; j++ {
				remove(cs, key)
			}
		}(i)
	}

	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			key := "shared"
			for j := 0; j < iterations; j++ {
				_ = get(cs, key)
			}
		}(i)
	}

	wg.Wait()
	// 如果代码中锁使用正确，则不会发生数据竞争
	// 最终 map 状态可能为空或非空，但程序不应崩溃
}
