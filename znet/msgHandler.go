package znet

import (
	"fmt"
	"game-demo/utils"
	"game-demo/ziface"
	"strconv"
)

type MsgHandler struct {
	// 存放msgId和对应的处理方法
	Apis map[uint32]ziface.IRouter
	// 负责Worker的消息队列
	TaskQueue []chan ziface.IRequest
	// 业务工作的线程池数目
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
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

func (m *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerPoolSize)

		go func(workerIdx int) {
			fmt.Println("Worker ID=", workerIdx, " is started ! ")
			for {
				select {
				case request := <-m.TaskQueue[workerIdx]:
					m.DoMsgHandler(request)
				}
			}
		}(i)
	}
}

func (m *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	workerId := request.GetConnection().GetConnID() % m.WorkerPoolSize
	fmt.Println("Add connId = ", request.GetConnection().GetConnID(), " request MsgId= ", request.GetMsgId(), " to WorkerId", workerId)
	m.TaskQueue[workerId] <- request
}
