package user

import (
	"yftxhub/pkg/app"
	"yftxhub/pkg/database"
	"yftxhub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断 Phone 已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (Usermodel User) {
	database.DB.Model(User{}).Where("phone = ?", loginID).Or("email = ?", loginID).Or("name = ?", loginID).First(&Usermodel)
	return
}

// GetByPhone 通过 手机号 来获取用户
func GetByPhone(phone string) (Usermodel User) {
	database.DB.Model(User{}).Where("phone = ?", phone).First(&Usermodel)
	return
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (Usermodel User) {
	database.DB.Model(User{}).Where("email = ?", email).First(&Usermodel)
	return
}

// Pagiate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return users, paging
}
