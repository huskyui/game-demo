package znet

import (
	"errors"
	"fmt"
	"game-demo/utils"
	"game-demo/ziface"
	"io"
	"net"
)

// 链接模块
type Connection struct {
	TcpServer ziface.IServer

	Conn *net.TCPConn

	ConnID uint32

	IsClose bool

	ExitChan chan bool

	MsgHandler ziface.IMsgHandler

	msgChan chan []byte
}

func (c *Connection) StartWriter() {
	fmt.Println("writer Goroutine is running")
	defer fmt.Println("connID=", c.ConnID, " Writer is exit,remote addr is ", c.RemoteAddr().String())

	for {
		select {
		case msg := <-c.msgChan:
			_, err := c.Conn.Write(msg)
			if err != nil {
				fmt.Println("write err", err)
				return
			}
		case existFlag := <-c.ExitChan:
			if existFlag {
				fmt.Println("捕获exitChan")
				return
			}
		}
	}
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
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandler(&req)
		}

	}

}

func (c *Connection) Start() {
	fmt.Println("Connection start,connId", c.ConnID)
	// 启动当前业务读数据
	go c.StartReader()
	// 启动writer
	go c.StartWriter()
	// call conn start hook
	c.TcpServer.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	fmt.Println("Conn stop", c.ConnID)
	if c.IsClose {
		return
	}
	c.IsClose = true
	c.TcpServer.CallOnConnStop(c)
	c.Conn.Close()
	c.ExitChan <- true
	c.TcpServer.GetConnMgr().Remove(c)
	close(c.ExitChan)
	close(c.msgChan)

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
	c.msgChan <- pack
	return nil
}

// 初始化链接
func NewConnection(server ziface.IServer, conn *net.TCPConn, connId uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connId,
		IsClose:    false,
		ExitChan:   make(chan bool, 1),
		MsgHandler: msgHandler,
		msgChan:    make(chan []byte, 1),
	}
	c.TcpServer.GetConnMgr().Add(c)
	return c
}
