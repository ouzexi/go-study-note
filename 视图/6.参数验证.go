package main

import "github.com/gin-gonic/gin"

type SignUserInfo struct {
	Name string `json:"name" binding:"required,min=4,max=6,contains=f"` // 用户名
	Age  int    `json:"age" binding:"gt=18,lt=30"`                      // 年龄
	// Password   string `json:"password"`                               // 密码
	// RePassword string `json:"re_password" binding:"eqfield=Password"` // 确认密码
	Sex      string   `json:"sex" binding:"oneof=man woman"`
	LikeList []string `json:"like_list" binding:"required,dive,required,startswith=like"`
	// 对于数组来说，dive后面的验证就是针对数组中的每一个元素
	IP   string `json:"ip" binding:"ip"`
	Url  string `json:"url" binding:"url"`
	Uri  string `json:"uri" binding:"uri"`
	Date string `json:"date" binding:"datetime=2006-01-02 15:04:06"`
	// 1月2号下午3点4分5秒在06年
}

func main() {
	router := gin.Default()

	router.POST("/", func(ctx *gin.Context) {

		var user SignUserInfo

		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": user})
	})

	router.Run(":80")
}
