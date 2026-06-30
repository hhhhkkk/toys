package consistencyhash

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

// 1. 基础功能测试：添加节点后，环的长度是否正确
func TestAddNode(t *testing.T) {
	h := New() // 每个节点 5 个虚拟节点
	h.AddNode(Config{
		NodeName: "server1",
		Replica:  5,
	})

	if len(h.ring) != 5 {
		t.Errorf("ring length = %d, expected %d", len(h.ring), 5)
	}

	h.AddNode(Config{
		NodeName: "server2",
		Replica:  5,
	})
	if len(h.ring) != 10 {
		t.Errorf("ring length = %d, expected %d", len(h.ring), 10)
	}
}

// 2. 核心功能测试：相同的 Key 永远指向同一个节点（确定性）
func TestGetNodeConsistency(t *testing.T) {
	// h := New(10)
	// h.AddNode("A")
	// h.AddNode("B")
	// h.AddNode("C")
	h := New()
	h.AddNode(Config{
		NodeName: "A",
		Replica:  10,
	})
	h.AddNode(Config{
		NodeName: "B",
		Replica:  10,
	})
	h.AddNode(Config{
		NodeName: "C",
		Replica:  10,
	})

	key := "user_12345"
	node1 := h.Get(key)
	node2 := h.Get(key)

	if node1 != node2 {
		t.Errorf("Consistency failed: same key got %s and %s", node1, node2)
	}
	if node1 == "" {
		t.Error("Get returned empty string for existing key")
	}
}

// 3. 均匀性测试（负载均衡）：验证 10000 个 Key 的分布偏差 < 20%
func TestDistribution(t *testing.T) {
	h := New() // 虚拟节点多一点，分布更均匀
	nodes := []Config{Config{NodeName: "Node-A", Replica: 160}, Config{NodeName: "Node-B", Replica: 160}, Config{NodeName: "Node-C", Replica: 160}}
	for _, n := range nodes {
		h.AddNode(n)
	}

	// 生成 10000 个测试 Key
	keyCount := 10000
	counts := make(map[string]int)

	for i := 0; i < keyCount; i++ {
		key := "key_" + strconv.Itoa(i)
		target := h.Get(key)
		counts[target]++
	}

	// 期望平均分布：10000 / 3 ≈ 3333
	expected := keyCount / len(nodes)
	tolerance := float64(expected) * 0.2 // 允许 20% 波动

	for _, node := range nodes {
		actual := counts[node.NodeName]
		diff := float64(actual - expected)
		if diff < 0 {
			diff = -diff
		}
		if diff > tolerance {
			t.Errorf("Distribution imbalance: node %s got %d, expected ~%d (diff %.2f%%)",
				node.NodeName, actual, expected, (diff/float64(expected))*100)
		}
	}
}

// 4. 添加新节点后，数据迁移量应低于 1/N（一致性哈希的核心优势）
func TestMigrationRate(t *testing.T) {
	// h := New(50)
	// h.AddNode("Old1")
	// h.AddNode("Old2")

	h := New()
	h.AddNode(Config{
		NodeName: "Old1",
		Replica:  100,
	})
	h.AddNode(Config{
		NodeName: "Old2",
		Replica:  100,
	})

	// 记录旧环下 1000 个 Key 的归属
	oldMap := make(map[string]string)
	keys := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		k := "migrate_" + strconv.Itoa(i)
		keys[i] = k
		oldMap[k] = h.Get(k)
	}

	// 新增一台机器
	// h.AddNode("New3")
	h.AddNode(Config{
		NodeName: "New3",
		Replica:  100,
	})

	// 计算需要搬家的比例
	moved := 0
	for _, k := range keys {
		if newTarget := h.Get(k); newTarget != oldMap[k] {
			moved++
		}
	}

	movedRate := float64(moved) / float64(len(keys))
	// 理论上加 1 台机器，最多迁移 1/3 的数据（实际上远小于这个值）
	if movedRate > 0.4 { // 40% 作为一个宽松的阈值
		t.Errorf("Too much data migrated: %.2f%%, expected < 40%%", movedRate*100)
	}
	t.Logf("Migration rate after adding 1 node: %.2f%%", movedRate*100)
}

// 5. 并发安全测试（必须启用 race detector）
func TestConcurrentAccess(t *testing.T) {
	h := New()

	var wg sync.WaitGroup
	wg.Add(20)

	// 10 个写协程（添加节点）
	for i := 0; i < 10; i++ {
		go func(idx int) {
			defer wg.Done()
			h.AddNode(
				Config{
					NodeName: fmt.Sprintf("writer_%d", idx),
					Replica:  3,
				})
		}(i)
	}

	// 10 个读协程（获取 Key）
	for i := 0; i < 10; i++ {
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("read_%d_%d", idx, j)
				_ = h.Get(key) // 只读，不 panic 就算通过
			}
		}(i)
	}

	wg.Wait()
	// 如果代码没有加锁，运行 go test -race 时会立即报错
}

// 6. 边缘测试：空环
func TestEmptyRing(t *testing.T) {
	h := New()
	// 没有添加任何节点
	result := h.Get("any_key")
	if result != "" {
		t.Errorf("Expected empty string for empty ring, got %s", result)
	}
}
