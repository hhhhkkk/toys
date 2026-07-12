package main

import (
	"fmt"
	"sync"
	"time"
)

func demoCond() {
	var lock sync.Mutex

	var q []int

	cond := sync.Cond{L: &lock}

	for i := range 3 {
		go func(i int) {
			for {
				// 锁了所有？
				cond.L.Lock()

				for len(q) == 0 {
					cond.Wait()
					fmt.Printf("%dwait\n", i)
				}
				j := q[0]
				q = q[1:]
				// 干啥?
				cond.L.Unlock()
				fmt.Printf("%d 消费%d\n", i, j)
				time.Sleep(2 * time.Second)
			}
		}(i)
	}

	go func() {
		for i := 0; ; i++ {
			cond.L.Lock()

			q = append(q, i)
			// 唤醒一个
			cond.L.Unlock()
			cond.Signal()
			time.Sleep(1 * time.Second)
		}

	}()

	select {}
}
