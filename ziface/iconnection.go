package ziface

import "net"

type IConnection interface {

	// 启动链接，让当前链接准备开始工作
	Start()
	// 停止链接，解释当前链接
	Stop()
	// 获取当前链接的socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前链接模块的链接id
	GetConnID() uint32
	// 获取远程客户端的TCP状态 IP port
	RemoteAddr() net.Addr
	// 发送数据 将数据发送给客户端
	Send(msgId uint32, data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
