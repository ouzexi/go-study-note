package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 中间件m1
func m1(c *gin.Context) {
	fmt.Println("m1")
	c.JSON(200, gin.H{"msg": "m1的响应"})
	// 拦截 - 中断后面的中间件
	c.Abort()
}

func m2(c *gin.Context) {
	fmt.Println("m2 in")
	c.Next()
	fmt.Println("m2 out")
}

func main() {
	router := gin.Default()

	router.GET("/", m1, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "index"})
	}, func(ctx *gin.Context) {
		fmt.Println(1)
	}, func(ctx *gin.Context) {
		fmt.Println(2)
	}, func(ctx *gin.Context) {
		fmt.Println(3)
	})

	router.Run(":80")
}
