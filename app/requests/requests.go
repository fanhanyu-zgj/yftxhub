// Package requests 处理请求数据和表单验证
package requests

import (
	"yftxhub/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

// Validate
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// 1. 解析请求，支持 JSON 数据，表单请求和 URL Query
	if err := c.ShouldBind(obj); err != nil {
		// 解析失败，返回 402 状态码和错误信息
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式")
		return false
	}
	// 2. 验证表单
	errs := handler(obj, c)
	// 3. 判断验证是否通过
	if len(errs) > 0 {
		response.ValidationError(c, errs)
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

func validateFile(c *gin.Context, data interface{}, rules govalidator.MapData, message govalidator.MapData) map[string][]string {

	// 配置初始化
	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      message,
	}

	// 开始开始验证
	return govalidator.New(opts).Validate()
}
