package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/ddshop/user_srv/cmd/api/handler"
)

func NewRouter() *gin.Engine {
	
	router := gin.New()

	// 注册业务路由 
	registerApiRoutes(router)

	// 配置 404 路由
	setupNoFoundHandler(router)

	return router
} 


// 注册业务路由
func registerApiRoutes(router *gin.Engine) {
	r := router

	v1 := r.Group("/v1")

	// 验证码服务
	v1.GET("/verify", handler.ShowVerifyCode)

	// 用户服务
	user := v1.Group("/user")
	user.Use()

	{
		user.POST("/login", handler.PasswordLogin)	
		user.POST("/register", handler.Register)
		user.GET("/list", handler.JWTAuth, handler.GetUserList)		
	}
	

	// {
	// 	note1.GET("/query", handlers.QueryNote)
	// 	note1.POST("", handlers.CreateNote)
	// 	note1.PUT("/:note_id", handlers.UpdateNote)
	// 	note1.DELETE("/:note_id", handlers.DeleteNote)		
	// }

}


func setupNoFoundHandler(router *gin.Engine) {
	
	// 处理错误路由，精确匹配
	router.NoRoute(func(ctx *gin.Context) {

		// 获取 header 里面的 'Accept' 信息
		acceptStr := ctx.Request.Header.Get("Accept")

		if strings.Contains(acceptStr, "text/html") {

			// 如果是 HTML
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {

			// 默认返回 JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确",
			})
		}
	})
}

