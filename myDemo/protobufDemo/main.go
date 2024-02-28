package main

import (
	"fmt"
	"game-demo/myDemo/protobufDemo/pb"
	"google.golang.org/protobuf/proto"
)

func main() {

	person := &pb.Person{
		Name:   "xxx",
		Age:    16,
		Emails: []string{"qq.com", "baidu.com"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "110",
				Type:   pb.PhoneType_MOBILE,
			},
			&pb.PhoneNumber{
				Number: "120",
				Type:   pb.PhoneType_WORK,
			},
		},
	}
	data, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("marshal err", err)
		return
	}

	newData := &pb.Person{}
	err = proto.Unmarshal(data, newData)
	if err != nil {
		fmt.Println("unmarshal err", err)
	}
	fmt.Println("源数据", data)
	fmt.Println("解码之后的数据", newData)

}
