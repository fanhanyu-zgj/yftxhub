// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "yftxhub/app/http/controllers/api/v1"
	"yftxhub/app/models/user"
	"yftxhub/app/requests"
	"yftxhub/pkg/jwt"
	"yftxhub/pkg/response"

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
	response.JSON(c, gin.H{"exist": user.IsPhoneExist(request.Phone)})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 请求对象
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}
	response.JSON(c, gin.H{"exist": user.IsEmailExist(request.Email)})
}

// SignupUsingPhone 使用手机和验证码进行注册
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	// 1.验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}
	// 2. 验证成功，创建数据
	userModel := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	userModel.Create()
	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.CreatedJSON(c, gin.H{
			"data":  userModel,
			"token": token,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}

// SignupUsingPhone 使用手机和验证码进行注册
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	// 1.验证表单
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}
	// 2. 验证成功，创建数据
	userModel := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	userModel.Create()
	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)

		response.CreatedJSON(c, gin.H{
			"data":  userModel,
			"token": token,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
