package bizjobs

import (
	"context"
	"fmt"
	"time"
)

type Test string

func (t Test) Name() string {
	return string(t)
}

func (t Test) Run(ctx context.Context) error {
	tick := time.Tick(10 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-tick:
			fmt.Println("service normal")
		}
	}
}
