package main

import "github.com/hhhhkkk/mini-blog/v1/internal"

func main() {
	app, err := internal.InitApp()
	if err != nil {
		panic("服务生成失败...")
	}
	if err := app.Run(); err != nil {
		panic("start error")
	}
}
