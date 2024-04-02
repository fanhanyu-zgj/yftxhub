package cache

import (
	"time"
	"yftxhub/pkg/config"
	"yftxhub/pkg/redis"
)

// RedisStore 实现 cache.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	rs := &RedisStore{}
	rs.RedisClient = redis.NewClient(address, username, password, db)
	rs.KeyPrefix = config.Get("app.name") + ":cache:"
	return rs
}

func (s *RedisStore) Set(key string, value string, expireTime time.Duration) {
	s.RedisClient.Set(key+s.KeyPrefix, value, expireTime)
}

func (s *RedisStore) Get(key string) string {
	return s.RedisClient.Get(key + s.KeyPrefix)
}

func (s *RedisStore) Has(key string) bool {
	return s.RedisClient.Has(key + s.KeyPrefix)
}

func (s *RedisStore) Forget(key string) {
	s.RedisClient.Del(key + s.KeyPrefix)
}

func (s *RedisStore) Forever(key string, value string) {
	s.RedisClient.Set(key+s.KeyPrefix, value, 0)
}

func (s *RedisStore) Flush() {
	s.RedisClient.FlushDB()
}

func (s *RedisStore) Increment(parameters ...interface{}) {
	s.RedisClient.Increment(parameters)
}

func (s *RedisStore) Decrement(parameters ...interface{}) {
	s.RedisClient.Decrement(parameters)
}

func (s *RedisStore) IsAlive() error {
	return s.RedisClient.Ping()
}
