package main

import (
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.String(200, "Hello World")
}

func main() {
	// 创建一个默认的路由
	router := gin.Default()

	// 绑定路由规则和路由函数，访问/index的路由，将由对应的函数处理
	router.GET("/index", Index)

	// 启动监控,gin会把web服务运行在本机的0.0.0.0的8080端口上
	// router.Run(":8080")

	// 修改ip为内网ip
	router.Run("0.0.0.0:8080")

	// 用原生http服务的方式,router.Run本质就是http.ListenAndServe的进一步封装
	// http.ListenAndServe(":8080", router)
}
