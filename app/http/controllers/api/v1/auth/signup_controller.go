// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"net/http"
	v1 "yftxhub/app/http/controllers/api/v1"
	"yftxhub/app/models/user"
	"yftxhub/app/requests"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 请求对象
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}
	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(request.Phone)})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 请求对象
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{"exist": user.IsEmailExist(request.Email)})
}
