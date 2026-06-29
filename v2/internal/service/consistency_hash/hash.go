package consistencyhash

import (
	"fmt"
	"hash/crc32"
	"sort"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

type MyHash struct {
	hash     Hash
	ring     []int
	nodeMap  map[int]string
	Replicas uint
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
	for i := 0; i < int(h.Replicas); i++ {
		key := fmt.Sprintf("node%s+%d", nodeName, i)
		index := int(h.hash([]byte(key)))
		h.nodeMap[index] = nodeName
		h.ring = append(h.ring, index)
	}

	sort.Ints(h.ring)
}
