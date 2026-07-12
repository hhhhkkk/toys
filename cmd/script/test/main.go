package main

import "time"

func main() {
	var n = 123

	c := make(chan int)

	go func() {
		c <- n + 0
	}()

	time.Sleep(time.Second)
	n = 789
	println(<-c)
}
