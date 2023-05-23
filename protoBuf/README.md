
## 快速开始
1. 安装 protoc

2. 编写helloword.proto文件
定义 message, service


3. 安装 `protoc-gen-go`, 生成 message代码

```
go install github.com/golang/protobuf/protoc-gen-go

protoc  --proto_path=. --go_out=./ *.proto
```
生成了go 模型定义代码: `helloword.pb.go`





4. 安装 `protoc-gen-go`,生成 service代码
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --proto_path=. --go-grpc_out=./ *.proto
```
生成了grpc 服务定义代码 `helloword_grpc.pb.go`

### oneof 
一个oneof里可包括多种字段, 并且最多同时存在一个字段

## 使用生成的代码
[test.go](./test.go)

