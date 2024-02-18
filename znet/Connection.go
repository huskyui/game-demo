package znet

import (
	"fmt"
	"game-demo/ziface"
	"net"
)

// 链接模块
type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	IsClose bool

	handleApi ziface.HandleFunc

	ExitChan chan bool
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println("connID=", c.ConnID, " Reader is exit,remote addr is ", c.RemoteAddr().String())
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}
		// 调用当前链接所绑定的HandleAPI
		if err := c.handleApi(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnId", c.ConnID, " handle is error", err)
			break
		}
	}

}

func (c *Connection) Start() {
	fmt.Println("Connection start,connId", c.ConnID)
	// 启动当前业务读数据
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Conn stop", c.ConnID)
	if c.IsClose {
		return
	}
	c.IsClose = true
	c.Conn.Close()

}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	//TODO implement me
	panic("implement me")
}

// 初始化链接
func NewConnection(conn *net.TCPConn, connId uint32, callBack ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connId,
		handleApi: callBack,
		IsClose:   false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
