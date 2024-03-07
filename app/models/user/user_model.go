// Package user 存放用户 Model 相关逻辑
package user

import (
	"yftxhub/app/models"
	"yftxhub/pkg/database"
	"yftxhub/pkg/hash"
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

// Create 创建用户，通过 User.ID 来判断是否成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparPassword 密码是否正确
func (userModel *User) ComparPassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}
