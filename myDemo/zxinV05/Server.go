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
	fmt.Println("call Router handle...")
	id := request.GetMsgId()
	fmt.Println("获取到数据msgId:", id, "数据:", string(request.GetData()))
	request.GetConnection().Send(id, []byte("获取到数据啦！"+string(request.GetData())))
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("after handle")
}

func main() {
	// 1.创建server
	s := znet.NewServer("[zxinV0.5]")
	s.AddRouter(&PingRouter{})

	// 2.运行
	s.Serve()

}
