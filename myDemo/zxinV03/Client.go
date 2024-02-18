package main

import (
	"fmt"
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
		_, err := conn.Write([]byte("hello zxinv0.1"))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read conn err", err)
			return
		}
		fmt.Printf("server callback %s,cnt:%d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}

}
