package main

import (
	"fmt"
	"game-demo/znet"
	"io"
	"net"
	"time"
)

/*
 */
func main() {
	fmt.Println("client start")
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err,exit")
		return
	}
	for {
		dp := znet.NewDataPack()
		msg := znet.NewMessagePack(1, []byte("Hello!"))

		pack, err := dp.Pack(msg)
		if err != nil {
			fmt.Println("pack error", err)
			break
		}

		_, err = conn.Write(pack)
		if err != nil {
			fmt.Println("write error", err)
			break
		}

		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("read head err", err)
			break
		}
		unPack, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack err", err)
			break
		}

		binaryData := make([]byte, unPack.GetMsgLen())
		_, err = io.ReadFull(conn, binaryData)
		if err != nil {
			fmt.Println("read binaryData error")
			break
		}

		fmt.Println("读到server的数据，msgId:", unPack.GetMsgId(), "数据：", string(binaryData))
		break
	}

}
