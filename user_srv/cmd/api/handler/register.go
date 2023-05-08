package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/sjxiang/ddshop/user_srv/cmd/api/requests"
	"github.com/sjxiang/ddshop/user_srv/pb"
)


func Register(ctx *gin.Context) {
	req := new(requests.RegisterUsingEmail)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
		})
		return
	}

	password, _ := strconv.Atoi(req.Password)

	// 拨号连接 gRPC 服务器
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 50051), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】", "Msg", err.Error())
	}
	
	// 生成 gRPC 的 client，并调用接口
	userSrvService := pb.NewUserClient(conn)

	// 注册逻辑
	res, err := userSrvService.CreateUser(context.Background(), &pb.CreateUserRequest{
		Mobile:   req.Mobile,
		Nickname: req.NickName,
		Email:    req.Email,
		Password: uint64(password),
	})

	if err != nil {
		zap.S().Errorw("[Register] 查询 【新建用户】失败")
		HandleGrpcErrorToHTTP(err, ctx)
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": res.Nickname,
	})


}