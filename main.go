package main

import (
	"flag"
	"fmt"
	"gohub/bootstrap"
	configBts "gohub/config"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载config目录下的配置信息
	configBts.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
	// r := gin.Default()

	// new 一个 Gin Engine 实例
	r := gin.New()

	bootstrap.SetupRoute(r)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
