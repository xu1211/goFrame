package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/xu1211/goFrame/protoBuf/helloProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to hello")
)

func main() {
	flag.Parse()
	// 连接 server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHellowordClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name, Password: 123})
	if err != nil {
		log.Fatalf("could not SayHello: %v", err)
	}
	log.Printf("调用SayHello服务返回: %s", r.GetMessage())
}
