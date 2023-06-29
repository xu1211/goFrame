[toc]

# go框架

## [gin(web框架)](./gin/README.md)


## [protoBuf 协议](./protoBuf/README.md)
protoBuf: 谷歌成熟的序列化结构化数据的开源机制

### grpc
默认使用protoBuf协议, 也可以使用其他协议 json等
- 引用proto 
`github.com/xu1211/goFrame/protoBuf`

- 服务端
[server/main.go](./grpc)
- 客户端
[client/main.go](./grpc/)
### micro

## [gorm (数据库)](./gorm)

- domain
  - test.go  (表对象struct)
  - test_CRUD.go  (封装sql操作)
- repository
  - repository.go  (数据库连接)
- main