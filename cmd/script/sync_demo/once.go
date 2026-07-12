package main

import (
	"fmt"
	"sync"
)

func main() {
	// 当前进程仅执行 1 次
	var once sync.Once
	for i := range 10 {
		once.Do(doOnce)
		// doOnce()
		fmt.Println(i)
	}
}

func doOnce() {
	fmt.Println("only once")
}
