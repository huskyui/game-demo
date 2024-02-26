package ziface

type IServer interface {
	Start()

	Stop()

	Serve()

	AddRouter(msgId uint32, router IRouter)

	GetConnMgr() IConnManager

	SetOnConnStart(func(connection IConnection))

	SetOnConnStop(func(connection IConnection))

	CallOnConnStart(conn IConnection)

	CallOnConnStop(conn IConnection)
}
