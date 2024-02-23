// Package user 存放用户 Model 相关逻辑
package user

import (
	"yftxhub/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitepmty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
