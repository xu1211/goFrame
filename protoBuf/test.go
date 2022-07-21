package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/xu1211/goFrame/protoBuf/helloProto"
	pboneof "github.com/xu1211/goFrame/protoBuf/oneof"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	helloword()
	fmt.Println("----------------------------------------------------------------")
	onefoTest()
}
func helloword() {
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

func onefoTest() {
	var (
		obj *pboneof.OneofTest
	)
	obj = &pboneof.OneofTest{
		OneofName: &pboneof.OneofTest_Filed1{Filed1: "string"},
	}
	fmt.Println(obj.String())
	fmt.Println("before Filed1: ", obj.GetFiled1())
	fmt.Println("before Filed2: ", obj.GetFiled2())
	fmt.Println()

	fmt.Println("------------json序列化   [外层有oneof字段]")
	jsonData, _ := json.Marshal(obj)
	fmt.Println(string(jsonData))
	fmt.Println("------------json反序列化 [失败]")
	obj2 := pboneof.OneofTest{}
	_ = json.Unmarshal(jsonData, &obj2)
	fmt.Println(obj2)
	fmt.Println("after one: ", obj2.GetFiled1())
	fmt.Println("after two: ", obj2.GetFiled2())
	fmt.Println()

	fmt.Println("------------protojson序列化  [外层无oneof字段]")
	protoJsonData, _ := protojson.Marshal(obj)
	fmt.Println(string(protoJsonData))
	fmt.Println("------------protojson反序列化 [成功]")
	obj2 = pboneof.OneofTest{}
	_ = protojson.Unmarshal(protoJsonData, &obj2)
	fmt.Println(obj2)
	fmt.Println("after one: ", obj2.GetFiled1())
	fmt.Println("after two: ", obj2.GetFiled2())
}
