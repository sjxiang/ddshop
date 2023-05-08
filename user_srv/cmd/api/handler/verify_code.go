package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/captcha"
)


// 验证码
func ShowVerifyCode(ctx *gin.Context) {
	id, b64s, err := captcha.New().Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误", 
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"captcha_id": id,
		"captcha_image": b64s,
	})
}
