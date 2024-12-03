package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/download", func(ctx *gin.Context) {
		// 需要唤起浏览器下载
		ctx.Header("Content-Type", "application/octet-stream")
		// 下载，设置文件名
		ctx.Header("Content-Disposition", "attachment; filename="+"牛逼.png")
		// 传输过程中的编码形式，乱码问题可能因为它
		ctx.Header("Content-Transfer-Encoding", "binary")
		// 这个只会浏览
		ctx.File("../uploads/13.png")
		// ctx.JSON(200, gin.H{"msg": "成功"})

		// 前后端分离-后端返回文件数据
	})

	router.Run(":80")
}
