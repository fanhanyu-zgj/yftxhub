package auth

import (
	v1 "yftxhub/app/http/controllers/api/v1"
	"yftxhub/app/models/user"
	"yftxhub/app/requests"
	"yftxhub/pkg/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseAPIController
}

func (ps *PasswordController) ResetByPhone(c *gin.Context) {
	// 1.验证表单
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	// 2. 更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}

}
