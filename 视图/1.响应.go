package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func _string(ctx *gin.Context) {
	ctx.String(http.StatusOK, "你好啊")
}

func _json(ctx *gin.Context) {
	// json响应结构体
	type UserInfo struct {
		UserName string `json:"username"`
		Age      int    `json:"age"`
		Password string `json:"-"` // 忽略转化为json
	}

	// user := UserInfo{"小小", 23, "123456"}

	// json响应map
	/* userMap := map[string]string{
		"username": "大大",
		"age":      "23",
	} */

	// 直接响应
	ctx.JSON(200, gin.H{"username": "嘎嘎", "age": 23})

	// ctx.JSON(200, user)
	// ctx.JSON(200, userMap)
}

func _html(ctx *gin.Context) {
	type UserInfo struct {
		UserName string
		Age      int
		Password string
	}
	// user := UserInfo{"彳亍", 12, "222"}

	ctx.HTML(200, "index.html", gin.H{"username": "枫枫"}) // 加载后台传递的数据，类似spring的modelAndView
	// ctx.HTML(200, "index.html", user) // 可以传递结构体
}

// 重定向
func _redirect(ctx *gin.Context) {
	ctx.Redirect(301, "http://www.baidu.com")
}

func main() {
	router := gin.Default()

	// 加载全局文件，加载目录下所有文件
	// 加载模板目录下所有文件
	router.LoadHTMLGlob("../templates/*")

	// 配置单个文件 - 第一个参数是网站访问的地址，第二个参数是文件所在位置
	// 在golang中，没有相对文件的路径，他只有相对项目的路径
	router.StaticFile("/test.jpg", "../static/test.jpg")

	// 配置静态目录,第二个参数是一个目录
	router.StaticFS("/static", http.Dir("../static/static"))

	router.GET("/", _string)
	router.GET("/json", _json)

	// 响应xml和yaml 略

	// 响应html
	router.GET("/html", _html)

	// 响应文件

	// 重定向
	router.GET("/redirect", _redirect)

	router.Run(":80")
}
