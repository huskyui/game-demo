package znet

import (
	"errors"
	"fmt"
	"game-demo/ziface"
	"io"
	"net"
)

// 链接模块
type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	IsClose bool

	ExitChan chan bool

	MsgHandler ziface.IMsgHandler
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println("connID=", c.ConnID, " Reader is exit,remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		dp := NewDataPack()
		headLen := dp.GetHeadLen()
		headData := make([]byte, headLen)
		_, err := io.ReadFull(c.Conn, headData)
		if err != nil {
			fmt.Println("read head data error", err)
			break
		}
		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack head data error", err)
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(c.Conn, data)
			if err != nil {
				fmt.Println("read full error")
				break
			}
			msg.SetData(data)
		}
		req := Request{
			connection: c,
			msg:        msg,
		}

		go c.MsgHandler.DoMsgHandler(&req)
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

func (c *Connection) Send(msgId uint32, data []byte) error {
	if c.IsClose == true {
		return errors.New("Connection closed when send msg")
	}
	dp := NewDataPack()
	msg := NewMessagePack(msgId, data)
	pack, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("pack error", err)
		return err
	}
	_, err = c.Conn.Write(pack)
	if err != nil {
		return err
	}
	return nil
}

// 初始化链接
func NewConnection(conn *net.TCPConn, connId uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connId,
		IsClose:    false,
		ExitChan:   make(chan bool, 1),
		MsgHandler: msgHandler,
	}
	return c
}
