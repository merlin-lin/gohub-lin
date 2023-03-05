package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"gohub/pkg/verifycode"

	"github.com/gin-gonic/gin"
)

type VerfifyCodeController struct {
	v1.BaseAPIController
}

func (vc *VerfifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()

	logger.LogIf(err)

	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// SendUsingPhone 发送手机验证码
func (vc *VerfifyCodeController) SendUsingPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// 2. 发送SMS
	if ok := verifycode.NewVerfiyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "短信发送失败")
	} else {
		response.Success(c)
	}
}

// SendUsingEmail 发送 Email 验证码
func (vc *VerfifyCodeController) SendUsingEmail(c *gin.Context) {
	// 1. 验证表单
	request := requests.VerfifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerfifyCodeEmail); !ok {
		return
	}

	// 2. 发送邮件
	errs := verifycode.NewVerfiyCode().SendEmail(request.Email)
	if errs != nil {
		response.Abort500(c, "发送验证码失败～")
	} else {
		response.Success(c)
	}

}
