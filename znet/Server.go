package znet

import (
	"encoding/json"
	"fmt"
	"game-demo/utils"
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

	MsgHandler ziface.IMsgHandler

	ConnMgr ziface.IConnManager

	OnConnStart func(conn ziface.IConnection)

	OnConnStop func(conn ziface.IConnection)
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("Add Router success!")
}

func (s *Server) Start() {
	s.MsgHandler.StartWorkerPool()
	marshal, _ := json.Marshal(s)
	fmt.Println("server config ", string(marshal))

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
			if s.ConnMgr.Len() > utils.GlobalObject.MaxConn {
				conn.Close()
				fmt.Println("已经达到最大值，拒绝连接")
				continue
			}

			dealConn := NewConnection(s, conn, cid, s.MsgHandler)

			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func (s *Server) Stop() {
	fmt.Println("[stop] zinx server name ", s.Name)

	s.ConnMgr.ClearAll()
}

func (s *Server) Serve() {
	// 启动server
	s.Start()
	// 阻塞

	select {}
}

func (s *Server) SetOnConnStart(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("------>CallOnConnStart--------->")
		s.OnConnStart(conn)
	}

}

func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("------------->CallOnConnStop------------>")
		s.OnConnStop(conn)
	}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandler(),
		ConnMgr:    NewConnManager(),
	}
}
