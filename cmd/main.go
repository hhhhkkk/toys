package main

import (
	"github.com/hhhhkkk/mini-blog/v1/cmd/attach"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	// v1
	g.Go(func() error {
		return attach.Run()
	})

	// v2
	// g.Go(func() error {})

	g.Wait()
}
