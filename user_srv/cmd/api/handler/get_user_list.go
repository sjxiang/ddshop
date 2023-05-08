package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/sjxiang/ddshop/user_srv/pb"
)


func GetUserList(ctx *gin.Context) {

	// 取出 Get 中的参数
	pageNum := ctx.DefaultQuery("pageNum", "0")
	pageSize := ctx.DefaultQuery("pageSize", "0")
	num, err := strconv.Atoi(pageNum)
	size, err := strconv.Atoi(pageSize)
	

	// 拨号连接 gRPC 服务器
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 50051), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】", "Msg", err.Error())
	}
	
	// 生成 gRPC 的 client，并调用接口
	userSrvService := pb.NewUserClient(conn)
	
	res, err := userSrvService.GetUserList(ctx, &pb.GetUserListRequest{
		PageInfo: &pb.GetUserListRequest_Pagination{
			PageNum: uint32(num),
			PageSize: uint32(size),
		},
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 【用户列表】失败")
		HandleGrpcErrorToHTTP(err, ctx)
		return 
	}

	
	data := make([]interface{}, 0)
	for _, val := range res.Data {
		tmp := make(map[string]interface{})

		tmp["name"] = val.Nickname
		tmp["mobile"] = val.Email
		tmp["birthday"] = val.Birthday

		data = append(data, tmp)
	}

	ctx.JSON(http.StatusOK, data)

} 