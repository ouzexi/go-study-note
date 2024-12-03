package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ArticleInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UserListView(c *gin.Context) {
	var userList = []UserInfo{
		{"xx", 23},
		{"yy", 24},
		{"zz", 25},
	}
	c.JSON(200, userList)
}

func ArticleListView(c *gin.Context) {
	var articleList = []ArticleInfo{
		{"xx", "23"},
		{"yy", "24"},
		{"zz", "25"},
	}
	c.JSON(200, articleList)
}

func UserRouterInit(router *gin.RouterGroup) {
	userManager := router.Group("user_manager")
	{
		userManager.GET("/user", UserListView)
	}
}

func main() {
	router := gin.Default()

	api := router.Group("/api")

	// 用户管理
	UserRouterInit(api)
	// 文章管理
	articleManager := api.Group("article_manager")
	{
		articleManager.GET("/article", ArticleListView)
	}

	router.Run(":80")
}
