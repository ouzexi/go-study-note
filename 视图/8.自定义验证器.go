package main

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 获取结构体中的msg参数
func _GetValidMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return ""
}

func signValid(fl validator.FieldLevel) bool {
	var nameList []string = []string{"凤凤", "张三", "fengfeng"}
	for _, nameStr := range nameList {
		name := fl.Field().Interface().(string)
		if name == nameStr {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()

	// 自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}

	router.POST("/", func(ctx *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required,sign" msg:"用户名校验失败"` //sign是自定义验证器
			Age  int    `json:"age" binding:"required" msg:"年龄校验失败"`
		}
		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {

			ctx.JSON(200, gin.H{"msg": _GetValidMsg(err, &user)})
			return
		}
		ctx.JSON(200, gin.H{"data": user})
	})

	router.Run(":80")
}
