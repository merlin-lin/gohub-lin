package config

import "gohub/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{

			// 应用名称
			"name": config.Env("APP_NAME", "Gohub"),

			"env": config.Env("APP_ENV", "production"),

			"debug": config.Env("APP_DEBUG", "false"),

			// 应用服务端口
			"port": config.Env("APP_PORT", "3000"),

			// 加密会话、JWT加密
			"key": config.Env("APP_KEY", ""),

			// 用以生成链接
			"url": config.Env("APP_URL", "http://localhost:3000"),

			// 设置时区， JWT里会使用， 日志记录里也会使用到
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
