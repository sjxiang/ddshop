package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/captcha"
	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/jwt"
	"github.com/sjxiang/ddshop/user_srv/cmd/api/requests"
	"github.com/sjxiang/ddshop/user_srv/pb"
)



func PasswordLogin(ctx *gin.Context) {
	// 表单验证
	passwordLogin := new(requests.PasswordLogin)
	if err := ctx.ShouldBindJSON(passwordLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
		})
		return
	}

	// 验证器
	if err := passwordLogin.Check(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "邮箱地址格式填写错误",
		})
		return
	}

	if ok := captcha.New().Verify(passwordLogin.CaptchaId, passwordLogin.Code); !ok {
		
		zap.S().Infow("验证码", passwordLogin.CaptchaId, passwordLogin.Code)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	// 拨号连接 gRPC 服务器
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 50051), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】", "Msg", err.Error())
	}
	
	// 生成 gRPC 的 client，并调用接口
	userSrvService := pb.NewUserClient(conn)

	// 登录逻辑
	res1, err := userSrvService.GetUserByEmail(context.Background(), &pb.GetUserByEmailRequest{
		Email: passwordLogin.Email,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"mobile": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "登录失败",
				})
			}
		}
		return
	}

	// 只是查询到有这个用户，还需进一步检查密码
	
	_, err = userSrvService.VerifyPassword(context.Background(), &pb.VerifyPasswordRequest{
		PlainText: passwordLogin.Password,
		PasswordHash: res1.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "密码错误",
		})
		return
	}

	// 生成 token
	token, _ := jwt.NewJWT().GenToken(int64(res1.Id), res1.Nickname, uint(res1.Role))

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
		"token": token,
	})





}	



