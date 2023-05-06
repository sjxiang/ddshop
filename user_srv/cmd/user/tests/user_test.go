package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sjxiang/ddshop/user_srv/pb"
)


var (
	userClient  pb.UserClient
	conn        *grpc.ClientConn
)

func Init() {
	var err error
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err = grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	
	// 新建 gRPC 客户端
	userClient = pb.NewUserClient(conn)
}

func TestGetUserList(t *testing.T) {
	Init()
	defer conn.Close()

	// 请求时间
	ctx , cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := userClient.GetUserList(ctx, &pb.GetUserListRequest{
									PageInfo: &pb.GetUserListRequest_Pagination{
										PageNum: 1,
										PageSize: 1,
									},
	})

	if err != nil {
		panic(err)
	}

	for _, user := range res.Data {
		t.Log(user.Nickname, user.Password)
	}
}
