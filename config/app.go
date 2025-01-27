// Package config
package config

import "yftxhub/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": config.Env("APP_NAME", "Yftxhub"),
			// 当前环境，用以区分多环境，一般为 loca, stage, production, test
			"env": config.Env("APP_ENV", "production"),
			// 是否进入调试模式
			"debug": config.Env("APP_DEBUG", false),
			// 应用服务端口
			"port": config.Env("APP_PORT", "3000"),
			// 加密会话, JWT 加密
			"key": config.Env("APP_KEY", "123456789"),
			// 用以生成链接
			"url": config.Env("APP_URL", "http://localhost:3000"),
			// 设置时区，JWT 里会使用，日志记录也会使用到
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),

			// API 域名，未设置的话所有 API URL 加 api 前缀，如 http://domain.com/api/v1/users
			"api_domain": config.Env("API_DOMAIN"),
		}
	})
}
