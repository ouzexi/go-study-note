package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 查询参数
func _query(ctx *gin.Context) {
	// user := ctx.Query("user")
	user, ok := ctx.GetQuery("user")
	users := ctx.QueryArray("user")
	userMap := ctx.QueryMap("user")

	fmt.Println(user)
	fmt.Println(ok)

	fmt.Println(users)
	fmt.Println(userMap)
}

// 动态参数
func _param(ctx *gin.Context) {
	fmt.Println(ctx.Param("user_id"))
	fmt.Println(ctx.Param("book_id"))
}

// 表单参数
func _form(ctx *gin.Context) {
	fmt.Println(ctx.PostForm("name"))
	fmt.Println(ctx.PostFormArray("name"))
	fmt.Println(ctx.DefaultPostForm("name", "四四")) // 用户没传使用的默认值
	form, err := ctx.MultipartForm()               // 接收所有的form参数 包括文件
	fmt.Println(form)
	fmt.Println(err)
}

// 封装解析json到结构体上的函数 - any类型接收指针不需要*
func bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	fmt.Println(string(body))
	fmt.Println(contentType)
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(obj)
	}
	return nil
}

// 原始参数
func _raw(c *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User
	// 这里 结构体对象要传指针
	err := bindJson(c, &user)
	fmt.Println(err)
	/* body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	fmt.Println(string(body))
	fmt.Println(contentType)
	switch contentType {
	case "application/json":
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		err := json.Unmarshal(body, &user)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(user)
	} */
}

func main() {
	router := gin.Default()

	router.GET("/query", _query)
	router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book/_id", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)

	router.Run(":8080")
}
