package consistencyhash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"

	"github.com/hhhhkkk/mini-blog/v2/pkg"
)

type Config struct {
	NodeName string
	Replica  int
}

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

type MyHash struct {
	hash    Hash
	ring    []int
	nodeMap map[int]string
	// Replicas uint
	mu sync.RWMutex
}

func New() *MyHash {
	return &MyHash{
		hash:    crc32.ChecksumIEEE,
		nodeMap: make(map[int]string),
		ring:    []int{},
		// Replicas: repeat,
	}
}

func (h *MyHash) List() map[string][]string {

	ret := make(map[string][]string)

	ret["rings"] = pkg.MapSlice(h.ring, func(i int) string { return strconv.Itoa(i) })

	ret["map"] = make([]string, 0)
	for k, v := range h.nodeMap {
		ret["map"] = append(ret["map"], fmt.Sprintf("%d - %s", k, v))
	}

	return ret
}

func (h *MyHash) AddNode(config Config) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for i := 0; i < config.Replica; i++ {
		key := fmt.Sprintf("node%s+%d", config.NodeName, i)
		index := int(h.hash([]byte(key)))
		h.nodeMap[index] = config.NodeName
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

func (h *MyHash) Remove(rn string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	// 直接重建 ring， 这样就不用删除了
	nRing := make([]int, 0)
	for idx, v := range h.ring {
		if h.nodeMap[v] != rn {
			nRing = append(nRing, v)
		} else {
			delete(h.nodeMap, idx)
		}
	}
	h.ring = nRing
}
