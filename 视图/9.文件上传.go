package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		// fmt.Println(file.Filename)
		// fmt.Println(file.Size / 1024) // 单位是字节
		// ctx.SaveUploadedFile(file, "../uploads/12.png")

		// 1 io.ReadAll 方法
		readerFile, _ := file.Open()
		// data, _ := io.ReadAll(readerFile)
		// fmt.Println(string(data))

		// 2 Create + Copy 方法
		writeFile, _ := os.Create("../uploads/13.png")
		defer writeFile.Close()
		n, _ := io.Copy(writeFile, readerFile)
		fmt.Println(n)

		// 3 上传多文件

		ctx.JSON(200, gin.H{"msg": "上传成功"})
	})

	router.POST("/uploads", func(ctx *gin.Context) {
		// 上传多文件
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			ctx.SaveUploadedFile(file, "../uploads/"+file.Filename)
		}
		ctx.JSON(200, gin.H{"msg": fmt.Sprintf("成功上传%d个文件", len(files))})
	})

	router.Run(":80")
}
