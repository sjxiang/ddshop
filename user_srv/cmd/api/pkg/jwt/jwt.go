package jwt

import (
	"errors"
	"time"

	jwtpkg "github.com/golang-jwt/jwt/v4"
	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/conf"
	"go.uber.org/zap"
)

// JWT 定义一个 jwt 对象
type JWT struct {

	// 密钥，用以加密 JWT，读取配置消息 APP_KEY
	SignKey []byte

	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration

}


func NewJWT() *JWT {

	refreshTime := 60  // 暂时用不上

	return &JWT{
		SignKey: []byte(conf.Conf.JWT.SecretKey),
		MaxRefresh: time.Duration(refreshTime) *time.Minute,
	}
}

const TokenExpireDuation = time.Hour * 720  // 24 x 30，1 个月

type JWTCustomClaims struct {
	ID          int64
	NickName    string
	AuthorityId uint

	jwtpkg.StandardClaims
}


// GenToken 生成 JWT
func (jwt *JWT) GenToken(userID int64, username string, authorityId uint) (string, error) {

	// 创建自定义的载体
	payload := JWTCustomClaims{
		ID: userID,
		NickName: username,
		AuthorityId: authorityId,

		StandardClaims: jwtpkg.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuation).Unix(),  // 过期时间
			// Issuer: settings.Conf.AppConfig.Name,                  // 签发者
        },
	}

	// 生成 token （对 payload 使用指定的签名方法， 生成签名对象）
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, payload)

	// 签名（密钥 + 签名对象，得到完整编码的字符串 token）
	tokenString, err := token.SignedString(jwt.SignKey)
	if err != nil {
		zap.L().Error("签名失败_jwt_GenToken", zap.Error(err))
		return "", errors.New("签名失败")
	}

	return tokenString, nil
}



// ParseToken 解析 JWT
func  (jwt *JWT) ParseToken(tokenString string) (*JWTCustomClaims, error) {
	
	// 1. 解码 
	// 参数 1. tokenString
	// 参数 2. payload 结构体指针 
	// 参数 3. 解析私钥的回调函数
	token, err := jwtpkg.ParseWithClaims(
		tokenString, 
		&JWTCustomClaims{}, 
		func(token *jwtpkg.Token) (interface{}, error) {
			return jwt.SignKey, nil
		},
	)

	// 2. 解析出错
	if err != nil {
		zap.L().Error("解析出错_jwt_ ParseToke", zap.Error(err))
		return nil, errors.New("解析出错")
	}


	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验

	// 断言
	claims, ok := token.Claims.(*JWTCustomClaims)

	// 验证
	if ok && token.Valid {
		return claims, nil 
	}

	return nil, errors.New("请求令牌无效")
}