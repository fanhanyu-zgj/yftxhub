// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"net/http"
	"strings"
	"yftxhub/app/http/middlewares"
	"yftxhub/routes"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleware(router)
	// 注册 API 路由
	routes.RegisterAPIRoutes(router)
	// 配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(gin.Recovery(),
		//gin.Logger(),
		//gin.Recovery(),
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		accaptString := c.Request.Header.Get("Accept")
		if strings.Contains(accaptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusOK, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusOK, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
