// 反射

package main

import (
	"fmt"
	"reflect"
)

type Userrr struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Printtt(inter interface{}) {
	switch x := inter.(type) {
	case Userrr:
		// x.Name = "水水水" // 修改是无效的,因为你传指针进来，空接口断言无法获取，所以要使用反射
		fmt.Println(x.Name)
		fmt.Println(x.Age)
	default:
		fmt.Println("无匹配")
		break
	}
}

func FPrint(inter interface{}) {
	// t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	// fmt.Println(t.Kind()) // 查看底层信息 - struct
	// fmt.Println(t.Elem()) // 不传指针会报错 - It panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.
	// fmt.Println(v.Kind())
	// fmt.Println(v.Elem())
	// fmt.Println(v.Interface()) // 获取接口

	/* for i := 0; i < v.NumField(); i++ {
		// 获取结构体字段的类型，键和值
		fmt.Println(t.Field(i).Type, t.Field(i).Name, v.Field(i), t.Field(i).Tag.Get("json")) // 获取tag的json字段值
	} */

	v.Elem().FieldByName("Name").SetString("HAHA")
	// fmt.Println(v.Field(0)) // 获取Name的值
	// fmt.Println(v.Field(1)) // 获取Age的值
}

func main() {
	var us = Userrr{
		Name: "SHSH",
		Age:  12,
	}
	// 修改前
	fmt.Println(us)
	// Printtt(us)
	FPrint(&us) // 不传指针Elem会报错,传指针获取Elem，就可以根据反射的键修改值
	// FPrint(&us)

	// 修改后
	fmt.Println(us)
}
