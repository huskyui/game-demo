package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"game-demo/utils"
	"game-demo/ziface"
)

type DataPack struct{}

func (d *DataPack) GetHeadLen() uint32 {
	// dataLen  uint32  +  id uint32 = 4+4 = 8
	return 8
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	// 写入id
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		return nil, err
	}
	// 写入dataLen
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		return nil, err
	}
	// 写入data
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

func (d *DataPack) UnPack(binaryData []byte) (ziface.IMessage, error) {
	dataBuff := bytes.NewReader(binaryData)
	msg := &Message{}

	err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id)
	if err != nil {
		return nil, err
	}
	// 读取dataLen
	err = binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}

	fmt.Println("msg", msg, "msgId", msg.GetMsgId(), "msgLen", msg.GetMsgLen())
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too Large msg package")
	}

	return msg, nil
}

func NewDataPack() *DataPack {
	return &DataPack{}
}
