package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

type data struct {
	value      int
	afterValue int
}

type Pipeline struct {
	fanoutch chan data
	faninch  chan data

	fanoutTimeout time.Duration

	fanoutCountChan chan int
}

func main() {
	p := &Pipeline{
		fanoutch:        make(chan data, 1000),
		faninch:         make(chan data, 100),
		fanoutTimeout:   1 * time.Second,
		fanoutCountChan: make(chan int),
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	// go FanoutHandler(ctx, p)
	go controlFanoutHandler(ctx, p)

	p.fanoutCountChan <- 10

	go producer(ctx, p)

	go FaninHandler(ctx, p)

	// 动态重启

	tick := time.NewTicker(3 * time.Second)
	defer tick.Stop()
	go func() {
		for range tick.C {
			i := rand.IntN(40)
			if i == 0 {
				continue
			}
			p.fanoutCountChan <- i
		}
	}()

	time.Sleep(100 * time.Second)
}

func do(d data) data {
	fmt.Println(d.value)
	n := rand.IntN(5)
	if n > 3 {
		panic("rand error")
	}
	time.Sleep(time.Duration(n) * time.Second)
	d.afterValue = d.value * 2
	return d
}

func FaninHandler(ctx context.Context, p *Pipeline) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done.")
			return
		case data, ok := <-p.faninch:
			if !ok {
				fmt.Println("fanin channel close")
				return
			}

			fmt.Printf("source value: %d, after value: %d\n", data.value, data.afterValue)
		}
	}
}

func FanoutHandler(ctx context.Context, p *Pipeline) {

	for {
		select {
		case data, ok := <-p.fanoutch:
			if !ok {
				fmt.Println("通道关闭")
				return
			}

			doneFlag := make(chan int, 1)
			nc, cancel := context.WithTimeout(ctx, p.fanoutTimeout)
			// defer cancel()
			go func() {
				defer func() {
					if err := recover(); err != nil {
						if e, ok := err.(error); ok {
							fmt.Println("手动err " + e.Error())
						} else {
							fmt.Println(err)
						}
						return
					}
				}()
				p.faninch <- do(data)
				doneFlag <- 1
			}()
			select {
			case <-nc.Done():
				fmt.Println("timeout")
			case <-doneFlag:
				fmt.Println("处理完成")
			}
			cancel()
		case <-ctx.Done():
			fmt.Println("控制 ctx cancel")
			return
		}
	}
}

func controlFanoutHandler(ctx context.Context, p *Pipeline) {

	var cancelFunc context.CancelFunc
	for {
		select {
		case <-ctx.Done():
			fmt.Println("主进程退出, 启动分销停止")
			return
		case num, ok := <-p.fanoutCountChan:
			fmt.Printf("启动 %d 个", num)
			if !ok {
				return
			}
			if cancelFunc != nil {
				cancelFunc()
			}
			cctx, cancel := context.WithCancel(ctx)
			cancelFunc = cancel
			for range num {
				go func() {
					FanoutHandler(cctx, p)
				}()
			}
		}
	}

}

func producer(ctx context.Context, p *Pipeline) {
	for i := range 100 * 1000 {
		select {
		case p.fanoutch <- data{value: i}:
		case <-ctx.Done():
			fmt.Println("主进程结束")
			return
		}
	}
}
