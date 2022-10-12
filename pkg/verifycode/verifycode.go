package verifycode

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/helpers"
	"gohub/pkg/logger"
	"gohub/pkg/redis"
	"gohub/pkg/sms"
	"strings"
	"sync"
)

type VerfiyCode struct {
	Store Store
}

var once sync.Once
var internalVerfiyCode *VerfiyCode

// NewVerifyCode 单例模式获取
func NewVerfiyCode() *VerfiyCode {
	once.Do(func() {
		internalVerfiyCode = &VerfiyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name"),
			},
		}
	})

	return internalVerfiyCode
}

// SendSMS 发送短信验证码，调用示例：
//         verifycode.NewVerifyCode().SendSMS(request.Phone)
func (vc *VerfiyCode) SendSMS(phone string) bool {
	// 生成验证码
	code := vc.generateVerfiyCode(phone)

	// 方便本地和API自动测试
	if !app.IsProduction() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}

	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": code},
	})
}

func (vc *VerfiyCode) CheckAnswer(key string, answer string) bool {
	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})

	// 方便开发，在非生产环境下，具备特殊前缀的手机号和 Email后缀，会直接验证成功
	if !app.IsProduction() && (strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) || strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix"))) {
		return true
	}

	return vc.Store.Verify(key, answer, false)
}

func (vc *VerfiyCode) generateVerfiyCode(key string) string {
	// 生成随机码
	code := helpers.RandomNumber(config.GetInt("verfiycode.code_length"))

	if !app.IsLocal() {
		code = config.GetString("verfiycode.debug_code")
	}

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

	vc.Store.Set(key, code)

	return code
}
