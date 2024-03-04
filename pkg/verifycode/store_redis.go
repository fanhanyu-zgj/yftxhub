package verifycode

import (
	"time"
	"yftxhub/pkg/app"
	"yftxhub/pkg/config"
	"yftxhub/pkg/redis"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 实现 verifycode.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	// 本地环境方便调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetFloat64("verifycode.debug_expire_time"))
	}
	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

// Set 实现 verifycode.Store interface 的 Set 方法
func (s *RedisStore) Get(key string, clear bool) (val string) {
	key = s.KeyPrefix + key
	val = s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// Set 实现 verifycode.Store interface 的 Set 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
