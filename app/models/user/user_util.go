package user

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	paginator "gohub/pkg/paginator"
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

// GetByEmail 通过Email获取用户
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

func Get(idStr string) (userMode User) {
	database.DB.Where("id", idStr).First(&userMode)
	return
}

func All() (users []User) {
	database.DB.Find(&users)
	return
}

func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
