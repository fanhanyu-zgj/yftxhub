// Package requests 处理请求数据和表单验证
package requests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

// Validate
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// 1. 解析请求，支持 JSON 数据，表单请求和 URL Query
	if err := c.ShouldBindJSON(&obj); err != nil {
		// 解析失败，返回 402 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式",
			"error":   err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return false
	}
	// 2. 验证表单
	errs := handler(obj, c)
	// 3. 判断验证是否通过
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求不通过，具体请查看 errors",
			"errors":  errs,
		})
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, message govalidator.MapData) map[string][]string {

	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      message,
	}

	// 开始开始验证
	return govalidator.New(opts).ValidateStruct()
}
