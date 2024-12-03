package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Test struct {
	Name string `json:"name"`
}

var user = Test{"haha"}

func m10(c *gin.Context) {
	fmt.Println("m10 ...in")
	// 传递值
	c.Set("name", "枫枫")
	c.Set("user", user)
	c.Next()
	fmt.Println("m10 ...out")
}

func main() {
	router := gin.Default()

	// 全局注册
	router.Use(m10)

	router.GET("/", func(ctx *gin.Context) {
		fmt.Println(ctx.Get("name"))
		_user, _ := ctx.Get("user")
		// 断言
		user := _user.(Test)
		fmt.Println(user.Name)
		ctx.JSON(200, gin.H{"msg": "index"})
	})

	router.Run(":80")
}
