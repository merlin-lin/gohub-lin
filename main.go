package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()

	// new 一个 Gin Engine 实例
	r := gin.New()

	bootstrap.SetupRoute(r)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
