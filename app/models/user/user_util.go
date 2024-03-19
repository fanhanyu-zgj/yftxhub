package user

import "yftxhub/pkg/database"

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
