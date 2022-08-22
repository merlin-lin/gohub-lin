package auth

import (
	"fmt"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}

	request := PhoneExistRequest{}

	// 解析JSON请求
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())

		return
	}

	// 从数据库查询
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
