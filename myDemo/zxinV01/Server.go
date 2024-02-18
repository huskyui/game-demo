package main

import "game-demo/znet"

func main() {
	// 1.创建server
	s := znet.NewServer("[zxinV0.1]")

	// 2.运行
	s.Serve()

}
