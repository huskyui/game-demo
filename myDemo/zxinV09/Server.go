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
	fmt.Println("call Ping Router handle...")
	id := request.GetMsgId()
	fmt.Println("获取到数据msgId:", id, "数据:", string(request.GetData()))
	request.GetConnection().Send(id, []byte("获取到数据啦！"+string(request.GetData())))
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("after handle")
}

type HelloRouter struct{}

func (h *HelloRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("before Handler")
}

func (h *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("call hello Router handle ....")
	id := request.GetMsgId()
	fmt.Println("获取到数据msgId:", id, "数据:", string(request.GetData()))
	request.GetConnection().Send(id, []byte("hello,获取到数据啦！"+string(request.GetData())))
}

func (h *HelloRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("postHandle")
}

func main() {
	// 1.创建server
	s := znet.NewServer("[zxinV0.5]")
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	// 2.运行
	s.Serve()

}
