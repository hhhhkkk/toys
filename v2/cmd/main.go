package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/v2/config"
	"github.com/hhhhkkk/mini-blog/v2/internal"
	consistencyhash "github.com/hhhhkkk/mini-blog/v2/internal/service/consistency_hash"
)

func main() {
	appConf := config.New()

	var wg sync.WaitGroup
	for _, v := range appConf.HostList {
		value := v
		wg.Add(1)
		go func(vv config.Host) {

			app, err := internal.InitApp()
			if err != nil {
				panic(fmt.Errorf("%s up failed", vv.Name))
			}

			parseHost := strings.Split(vv.Host, ":")
			portS := parseHost[1]
			port, _ := strconv.Atoi(portS)

			if port > 8000 {
				err = app.Run(port)
			} else {
				err = app.Run()
			}

			if err != nil {
				panic(fmt.Errorf("%s up failed, cause: %s", vv.Name, err.Error()))
			}
			defer wg.Done()
		}(value)
	}

	engine := gin.Default()

	hash := consistencyhash.New()
	for _, v := range appConf.HostList {
		hash.AddNode(consistencyhash.Config{
			NodeName: v.Name,
			Replica:  v.FakerNum,
		})
	}

	g := engine.Group("/cache")
	{
		g.GET("/:key", func(ctx *gin.Context) {
			key := ctx.Param("key")
			name := hash.Get(key)

			if name == "" {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "hash node is nil",
				})
				return
			}

			host := ""
			for _, v := range appConf.HostList {
				if v.Name == name {
					host = v.Host
					break
				}
			}

			client := http.Client{}
			resp, _ := client.Get(fmt.Sprintf("http://%s/app/cache/%s", host, key))
			value, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "call api err" + err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"value":  string(value),
				"node":   name,
				"key":    key,
				"hashes": hash.List(),
			})
		})

		g.POST("", func(ctx *gin.Context) {
			type AddDto struct {
				Key   string `form:"key" json:"key"`
				Value string `form:"value" json:"value"`
			}

			var dto AddDto
			if err := ctx.ShouldBindBodyWithJSON(&dto); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "params is invalid."})
				return
			}

			name := hash.Get(dto.Key)

			host := ""
			for _, v := range appConf.HostList {
				if v.Name == name {
					host = v.Host
					break
				}
			}

			jsonBody, err := json.Marshal(dto)
			if err != nil {
				ctx.JSON(http.StatusServiceUnavailable, gin.H{
					"msg": err.Error(),
				})
				return
			}

			client := http.Client{}
			h := fmt.Sprintf("http://%s/app/cache", host)
			fmt.Println(string(jsonBody))
			resp, err := client.Post(h, "application/json", bytes.NewReader(jsonBody))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "call api err" + err.Error(),
				})
				return
			}
			defer resp.Body.Close()

			bodyBytes, err := io.ReadAll(resp.Body)
			ctx.JSON(http.StatusOK, gin.H{
				"node":  name,
				"value": string(bodyBytes),
			})
		})
	}

	go func() {
		fmt.Println("main in http://localhost:8001")
		if err := engine.Run(":8001"); err != nil {
			panic("main server fail")
		}
	}()
	wg.Wait()
}
