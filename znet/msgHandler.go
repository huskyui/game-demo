package znet

import (
	"fmt"
	"game-demo/ziface"
	"strconv"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (m *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	msgId := request.GetMsgId()
	router, exists := m.Apis[msgId]
	if !exists {
		fmt.Println("尚未实现当前clientHandler")
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

func (m *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	_, ok := m.Apis[msgID]
	if ok {
		panic("repeat api,msgId = " + strconv.Itoa(int(msgID)))
	}
	m.Apis[msgID] = router

}
