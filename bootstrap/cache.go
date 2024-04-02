package bootstrap // 启动程序功能
import (
	"fmt"
	"yftxhub/pkg/cache"
	"yftxhub/pkg/config"
)

func SetupCache() {
	// 初始化缓存使用的 redis client ，使用专属缓存 DB
	rds := cache.NewRedisStore(fmt.Sprintf("%v:%v", config.Get("redis.host"), config.Get("redis.port")), config.Get("redis.username"), config.Get("redis.password"), config.GetInt("redis.database_cache"))
	cache.InitWithCacheStore(rds)
}
