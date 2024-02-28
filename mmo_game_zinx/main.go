package main

import "game-demo/znet"

func main() {
	s := znet.NewServer("MMO game server")

	// 创建和销毁时的钩子函数

	// 注册路由

	// 启动
	s.Serve()

}
