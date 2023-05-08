// 处理图片验证码逻辑

package captcha

import (

	"sync"

	"github.com/mojocn/base64Captcha"
	
	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/cache"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}


// once 确保 internalCaptcha 对象值初始化一次
var once sync.Once

// internalCaptcha 内部使用的 Captcha 对象
var internalCaptcha *Captcha


func New() *Captcha {
	once.Do(func() {

		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}

		// 使用全局 Redis 对象，并配置存储 key 的前缀
		store := RedisStore{
			RedisClient: cache.Redis,
			KeyPrefix: "Captchaget:",
		}

		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			80,   // 验证码图片高度
			240,  // 验证码图片宽度
			6,    // 验证码的长度 
			0.7,  // 数字的最大倾斜角度
			80,   // 图片背景里的混淆点数量
		)

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store) 
	})

	return internalCaptcha
}

// GenerateCaptcha 生成图片验证码
func (c *Captcha) Generate() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证码是否正确
func (c *Captcha) Verify(id string, answer string) (match bool) {

	// 第 3 个参数，是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单错误提交需要多次输入图片验证码
	return c.Base64Captcha.Verify(id, answer, false)
}