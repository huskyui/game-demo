package znet

import "game-demo/ziface"

type Request struct {
	connection ziface.IConnection
	msg        ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.connection
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
