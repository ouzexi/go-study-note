package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		// 首字母大小写不区分 单词与单词之间用-连接
		// 获取指定请求头
		// fmt.Println(ctx.GetHeader("User-Agent"))
		// 获取全部 - Header是一个普通的Map[string][]string
		// fmt.Println(ctx.Request.Header.Get("User-Agent"))
		fmt.Println(ctx.Request.Header)
		ctx.JSON(200, gin.H{"msg": "成功"})
	})

	// 爬虫和用户区别对待
	router.GET("/index", func(ctx *gin.Context) {
		userAgent := ctx.GetHeader("User-Agent")
		// 用正则匹配 或 字符串的包含匹配
		if strings.Contains(userAgent, "python") {
			// 爬虫来了
			ctx.JSON(0, gin.H{"data": "这是响应给爬虫的数据"})
			return
		}
		ctx.JSON(0, gin.H{"data": "这是响应给用户的数据"})
	})

	// 设置响应头
	router.GET("/res", func(ctx *gin.Context) {
		ctx.Header("Token", "ouzxxxxxxxxxxxxxxxx")
		ctx.Header("Content-Type", "application/text;charset=utf-8")
		ctx.JSON(200, gin.H{"data": "看看响应头"})
	})

	router.Run(":80")
}
