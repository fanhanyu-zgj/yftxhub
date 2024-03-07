// Package hash 哈希操作类
package hash

import (
	"yftxhub/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	// GenerateFromPassword 第二各参数是 cost 值。建议大于 12，数值越大耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)

	return string(bytes)
}

// BcryptCheck 对米明文密码和数据库的 hash 值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 判断字符串是否是哈希过的数据
// BcryptIsHashed
func BcryptIsHashed(str string) bool {
	// bcrypt 加密后的长度等于 60
	return len(str) == 60
}
