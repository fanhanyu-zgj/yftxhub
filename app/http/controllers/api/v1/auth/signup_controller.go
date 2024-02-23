// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"fmt"
	"net/http"
	v1 "yftxhub/app/http/controllers/api/v1"
	"yftxhub/app/models/user"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}
	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 402 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}
	c.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(request.Phone)})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 请求对象
	type EmailExistRequest struct {
		Email string `json:"email"`
	}
	request := EmailExistRequest{}
	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 402 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}
	c.JSON(http.StatusOK, gin.H{"exist": user.IsEmailExist(request.Email)})
}
