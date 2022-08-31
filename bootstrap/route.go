package bootstrap

import (
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleware(router)
	// 注册API路由
	routes.RegisterAPIRoutes(router)
	// 配置404路由
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 判断请求头Accept信息
	// html请求另外处理
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义， 请确认url和请求方法是否正确。",
			})
		}
	})
}
