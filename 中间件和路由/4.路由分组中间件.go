package main

import "github.com/gin-gonic/gin"

type _UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func _UserListView(c *gin.Context) {
	var userList = []_UserInfo{
		{"xx", 23},
		{"yy", 24},
		{"zz", 25},
	}
	c.JSON(200, userList)
}

func Middleware(msg string) gin.HandlerFunc {
	// 中间件加括号，返回值就需要符合HandlerFunc
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "123456" {
			c.Next()
			return
		}
		c.JSON(200, gin.H{"code": "1001", "msg": msg})
		c.Abort()
	}
}

func _UserRouterInit(router *gin.RouterGroup) {
	// 路由分组中间件
	// 中间件加括号，返回值就需要符合HandlerFunc
	// 中间件加括号的用处是可以传递参数如msg
	userManager := router.Group("user_manager").Use(Middleware("用户校验失败"))
	{
		userManager.GET("/user", _UserListView)
	}
}

func main() {
	router := gin.Default()
	// router := gin.New() // 用New也可以 但是就要手动加上日志等
	// router.Use(gin.Logger(), gin.Recovery()) // 类似default是new的默认配置

	api := router.Group("/api")

	api.GET("/login", func(ctx *gin.Context) {
		// panic("手动报错")
		ctx.JSON(200, gin.H{"data": "123456"})
	})

	// 用户管理
	_UserRouterInit(api)

	router.Run(":80")
}
