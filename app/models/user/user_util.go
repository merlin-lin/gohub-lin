package user

import (
	"gohub/pkg/database"
)

// IsEmailExist 判断email是否被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断phone是否被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginId string) (userModel User) {
	database.DB.Where("phone = ?", loginId).Or("email = ?", loginId).Or("name = ?", loginId).First(&userModel)
	return
}

// GetByPhone 通过手机号获取用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}