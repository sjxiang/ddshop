// 自定义存储驱动 包裹一层 Redis，重新实现一遍 base64Captcha.Store.interface

package captcha

import (
	"errors"
	"time"

	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/cache"
)

// 自定义存储驱动
// RedisStore 实现 base64Captcha.Store.interface
type RedisStore struct {
	RedisClient  *cache.RedisClient
	KeyPrefix    string
}



// Set 实现 base654Captcha.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) error {

	TTL := time.Minute * time.Duration(30)  // 30 Min

	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, TTL); !ok {
		return errors.New("无法存储图片验证码答案")
	}

	return nil 
}

// Get 实现 base654Captcha.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) string {

	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)

	if clear {
		s.RedisClient.Del(key)
	}

	return val
}


// Verify 实现 base654Captcha.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}