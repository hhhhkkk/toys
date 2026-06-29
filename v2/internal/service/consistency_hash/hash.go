package consistencyhash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

type MyHash struct {
	hash     Hash
	ring     []int
	nodeMap  map[int]string
	Replicas uint
	mu       sync.RWMutex
}

func New(repeat uint) *MyHash {
	return &MyHash{
		hash:     crc32.ChecksumIEEE,
		nodeMap:  make(map[int]string),
		ring:     []int{},
		Replicas: repeat,
	}
}

func (h *MyHash) AddNode(nodeName string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for i := 0; i < int(h.Replicas); i++ {
		key := fmt.Sprintf("node%s+%d", nodeName, i)
		index := int(h.hash([]byte(key)))
		h.nodeMap[index] = nodeName
		h.ring = append(h.ring, index)
	}

	sort.Ints(h.ring)
}

func (h *MyHash) Get(k string) string {
	hash := int(h.hash([]byte(k)))

	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.ring) == 0 {
		return ""
	}
	idx := sort.Search(len(h.ring), func(i int) bool {
		return h.ring[i] >= hash
	})

	idx = idx % len(h.ring)

	return h.nodeMap[h.ring[idx]]
}
