package main

import (
	v2 "github.com/hhhhkkk/mini-blog/v2/cmd"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	// v1
	// g.Go(func() error {
	// 	return attach.Run()
	// })

	// v2
	g.Go(func() error {
		return v2.Run()
	})

	g.Wait()
}
