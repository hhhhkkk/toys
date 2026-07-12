package main

import (
	"fmt"
	"sync"
)

func main() {

	multiSenderChan := make(chan int, 100)
	closeChan := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(1)

	for i := 0; i < 3; i++ {
		go func(i int) {
			for {
				select {
				case <-closeChan:
					return
				case multiSenderChan <- i + 1:
					// do something
					fmt.Printf("sender %d\n", i+1)
				}
			}
		}(i)
	}

	go func() {
		defer func() {
			close(closeChan)
			wg.Done()
		}()
		j := 0
		for i := range multiSenderChan {
			fmt.Printf("puller get %d\n", i)
			j += i
			if j > 1000 {
				return
			}
		}
	}()

	wg.Wait()
}
