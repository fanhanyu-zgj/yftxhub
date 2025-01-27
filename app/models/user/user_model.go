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

	Name string `json:"name,omitepmty"`

	City         string `json:"city,omitepmty"`
	Introduction string `json:"introduction,omitepmty"`
	Avatar       string `json:"avatar,omitepmty"`

	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

// Get 通过 ID 获取用户
func Get(idstr string) (userModel User) {
	database.DB.Where("id", idstr).First(&userModel)
	return
}

// All 获取所有用户数据
func All() (users []User) {
	database.DB.Find(&users)
	return users
}
