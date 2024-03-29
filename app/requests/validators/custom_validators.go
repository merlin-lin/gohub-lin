package validators

import (
	"gohub/pkg/captcha"
	"gohub/pkg/verifycode"
)

func ValidateCaptcha(captchaId, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaId, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}

	return errs
}

// ValidatePasswordConfirm 自定义规则，检查两次密码是否正确
func ValidatePasswordConfirm(password, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm{
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}

	return errs
}

// ValidateVerifyCode 自定义规则，验证『手机/邮箱验证码』
func ValidateVerifyCode(key, answer string, errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerfiyCode().CheckAnswer(key, answer);!ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}

	return errs
}