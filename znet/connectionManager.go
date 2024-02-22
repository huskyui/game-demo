package znet

import (
	"errors"
	"fmt"
	"game-demo/ziface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection

	connLock sync.Mutex
}

func (c *ConnManager) Add(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
	fmt.Println("connection add to Connection Manager successfully; conn Num = ", c.Len())
}

func (c *ConnManager) Remove(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, conn.GetConnID())
	fmt.Println("connection remove from connection Manager successfully;Conn num = ", c.Len())
}

func (c *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if conn, ok := c.connections[connId]; ok {
		return conn, nil
	}
	return nil, errors.New("Connection 未找到")
}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearAll() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	for connId, conn := range c.connections {
		conn.Stop()
		fmt.Println("connId  ", connId, "stop ")
		delete(c.connections, connId)
	}
	fmt.Println("clear all Connections successfully! Conn num", c.Len())
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}
