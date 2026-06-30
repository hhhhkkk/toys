package main

import (
	v1 "github.com/hhhhkkk/mini-blog/v1/cmd/attach"
	v2 "github.com/hhhhkkk/mini-blog/v2/cmd/attach"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	// v1
	g.Go(func() error {
		return v1.Run()
	})

	// v2
	g.Go(func() error {
		return v2.Run()
	})

	g.Wait()
}
