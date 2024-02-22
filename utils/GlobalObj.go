package utils

import (
	"game-demo/ziface"
)

type GlobalObj struct {
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	Version           string
	MaxConn           int
	MaxPackageSize    uint32
	WorkerPoolSize    uint32 // 当前worker工作词的大小
	MaxWorkerPoolSize uint32 // 允许用户开辟多少的worker
}

//  定义全局对象

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	//data, err := os.ReadFile("conf/zinx.json")
	//if err != nil {
	//	panic(err)
	//}
	//json.Unmarshal(data, &GlobalObject)
}

// 初始化GlobalObject
func init() {
	GlobalObject = &GlobalObj{
		Name:              "ZxinTcpServer",
		Version:           "V0.4",
		TcpPort:           8999,
		Host:              "0.0.0.0",
		MaxConn:           1000,
		MaxPackageSize:    4096,
		WorkerPoolSize:    8,
		MaxWorkerPoolSize: 1024,
	}
	GlobalObject.Reload()
}
