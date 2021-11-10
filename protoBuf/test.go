package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"goFrame/protoBuf/helloProto"
)

func main() {
	helloword := &helloProto.HelloRequest{
		Name:     "cosmo",
		Password: 123,
	}
	fmt.Println("原始数据" + helloword.String())

	// proto编码
	data, err := proto.Marshal(helloword)
	if err != nil {
		fmt.Println("编码失败")
	} else {
		fmt.Println("编码成功：")
		fmt.Println(data)
	}

	// proto解码
	decodeData := &helloProto.HelloRequest{}
	err = proto.Unmarshal(data, decodeData)
	if err != nil {
		fmt.Println("解码失败")
	} else {
		fmt.Println("解码成功：")
		fmt.Println(decodeData)
		fmt.Println(decodeData.Name)
		fmt.Println(decodeData.Password)
	}
}
