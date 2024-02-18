package znet

import (
	"errors"
	"fmt"
	"game-demo/ziface"
	"net"
)

type Server struct {
	// 服务器名称
	Name string
	// 服务器的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn handle callBackToClient]...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("Write Back error", err)
		return errors.New("CallbackToClient error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[start server] at IP: %s,Port: %d", s.IP, s.Port)
	go func() {
		// 1.获取tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("reslove tcp addr error:", err)
			return
		}
		// 监听tcp的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		var cid uint32
		cid = 0
		fmt.Println("start zinx server succ", s.Name, "succ,Listening")
		// 阻塞等待客户端链接，处理客户端链接业务
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	// todo 将服务器的资源、状态  回收等等
}

func (s *Server) Serve() {
	// 启动server
	s.Start()
	// 阻塞

	select {}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}