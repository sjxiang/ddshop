package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sjxiang/ddshop/user_srv/pb"
)

func HandleGrpcErrorToHTTP(err error, ctx *gin.Context)  {
	// 将 gRPC 的 code 转换成 http 的状态码

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}

			return
		}
	}
}


func GetUserList(ctx *gin.Context) {

	// 取出 Get 中的参数
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")

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

	zap.S().Debug("获取用户列表页")


} 