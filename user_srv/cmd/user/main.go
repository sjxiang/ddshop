package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/sjxiang/ddshop/user_srv/cmd/user/conf"
	"github.com/sjxiang/ddshop/user_srv/cmd/user/dal"
	"github.com/sjxiang/ddshop/user_srv/cmd/user/service"
	"github.com/sjxiang/ddshop/user_srv/pb"
)


func Init() {
	conf.Init()
	dal.Init()
}

func main() {

	addr := flag.String("addr",  "localhost", "IP 地址")
	port := flag.Int("port", 50051, "端口")

	// 监听
	listener, err := net.Listen("tcp",  fmt.Sprintf("%s:%d", *addr, *port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	// 初始化
	Init()
		
	// 注册服务
	server := grpc.NewServer()
	pb.RegisterUserServer(server, new(service.UserServiceImpl))

	log.Printf("server listening at %v\n", listener.Addr())
	
	// 启动服务
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}

}


// go run main.go -h