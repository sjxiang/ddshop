package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/sjxiang/ddshop/inventory_srv/internal/dal"
	"github.com/sjxiang/ddshop/inventory_srv/internal/pkg/conf"
	"github.com/sjxiang/ddshop/inventory_srv/internal/service"
	"github.com/sjxiang/ddshop/inventory_srv/pb"
)
	
func Init() {
	conf.Init()
	dal.Init()
}


func main() {

	addr := flag.String("addr",  "localhost", "IP 地址")
	port := flag.Int("port", 50052, "端口")

	// 监听
	listener, err := net.Listen("tcp",  fmt.Sprintf("%s:%d", *addr, *port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	// 初始化
	Init()
		
	// 注册服务
	server := grpc.NewServer()
	pb.RegisterInventoryServer(server, new(service.InventoryServiceImpl))

	log.Printf("server listening at %v\n", listener.Addr())
	
	
	// 启动服务
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
	
}