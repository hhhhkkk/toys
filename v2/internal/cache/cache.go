package cache

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Node struct {
	Size  int
	Value string
}

type CacheService struct {
	lock sync.RWMutex
	data map[string]*Node
}

func NewCacheService() *CacheService {
	return &CacheService{
		lock: sync.RWMutex{},
		data: make(map[string]*Node),
	}
}

type AddDto struct {
	Key   string `form:"key" json:"key"`
	Value string `form:"value" json:"value"`
}

func (c *CacheService) Add(ctx *gin.Context) {
	dto := &AddDto{}
	if err := ctx.ShouldBindBodyWithJSON(dto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "params is invalid."})
		return
	}

	if dto.Key == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "key can't be empty string."})
		return
	}

	node := &Node{
		Value: dto.Value,
		Size:  len(dto.Value),
	}
	add(c, dto.Key, node)
	ctx.JSON(http.StatusNoContent, "")
}

type GetDto struct {
	Key string `form:"key" json:"key"`
}

func (c *CacheService) Get(ctx *gin.Context) {
	// dto := &GetDto{}
	// if err := ctx.ShouldBindBodyWithJSON(dto); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "params is invalid."})
	// 	return
	// }
	// if dto.Key == "" {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "key can't be empty string."})
	// 	return
	// }

	key := ctx.Param("key")
	if key == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "key can't be empty string."})
		return
	}
	node := get(c, key)
	if node == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "key is not match value."})
		return
	}
	ctx.JSON(http.StatusOK, node.Value)
}

type DelDto GetDto

func (c *CacheService) Del(ctx *gin.Context) {
	// dto := &DelDto{}
	// if err := ctx.ShouldBindBodyWithJSON(dto); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "params is invalid."})
	// 	return
	// }
	// if dto.Key == "" {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "key can't be empty string."})
	// 	return
	// }
	key := ctx.Param("key")
	if key == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "key can't be empty string."})
		return
	}
	remove(c, key)
	ctx.JSON(http.StatusNoContent, "")
}

func add(c *CacheService, k string, v *Node) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[k] = v
}

func get(c *CacheService, k string) *Node {
	c.lock.RLock()
	defer c.lock.RUnlock()
	node, exists := c.data[k]
	if !exists {
		return nil
	}
	nd := *node
	return &nd
}

func remove(c *CacheService, k string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.data, k)
}
