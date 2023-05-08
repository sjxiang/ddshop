package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/jwt"
)

// 基于 JWT 的认证中间件
func JWTAuth(ctx *gin.Context) {
	
	// 客户端携带 Token 有三种方式，具体取决于实际业务情况
	// 1. 放在请求 header
	// 2. 放在请求 body 
	// 3. 放在 url 中
	//
	// 方式一，具体做法：
	// token 放在 Header 的 Authorization 中，例如 "bearer xxx.xxx.xxx"
	
	authHeader := ctx.Request.Header.Get("Authorization")
	
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "请求 header 中 auth 为空",
		})
		ctx.Abort()
		return
	}

	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "请求 header 中，auth 格式有误",
		})
	
		ctx.Abort()
		return
	}	

	// parts[1] 是获取到的 tokenString，可以用 jwt 包方法解析
	claims, err := jwt.NewJWT().ParseToken(parts[1]) 
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "无效的 token",
		})

		ctx.Abort()
		return
	}

	// 检查 expired time （重新登录，申请 JWT）
	if float64(time.Now().Unix()) > float64(claims.ExpiresAt) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "过期的 token",
		})

		ctx.Abort()
		return
	}
	

	// TODO 校验
	
	// 将当前请求的 userID 信息保存到请求的上下文 ctx 中
	ctx.Set("username", claims.NickName)
	ctx.Next()  
	
	// 后续的 handler 可以用 ctx.Get("userID") 来获取当前请求的用户信息

}


// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Origin, Authorization")  
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return 
		}

		c.Next()
	}
}