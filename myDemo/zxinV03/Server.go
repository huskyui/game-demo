package main

import (
	"fmt"
	"game-demo/ziface"
	"game-demo/znet"
)

type PingRouter struct{}

func (p *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("before handle")
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	connection := request.GetConnection()
	connection.GetTCPConnection().Write([]byte("pong"))
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("after handle")
}

func main() {
	// 1.创建server
	s := znet.NewServer("[zxinV0.1]")
	s.AddRouter(&PingRouter{})

	// 2.运行
	s.Serve()

}
