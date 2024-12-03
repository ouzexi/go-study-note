package main

import "fmt"

type User1 struct {
	Name     string
	age      int
	password string
}

func main() {
	// var u1 User1 = User1{Name: "张三", age: 21, password: "123ss"}
	var u2 User1
	u2.Name = "lisi"
	var u1 User1 = User1{"张三", 21, "123ss"}
	fmt.Printf("u1:%#v, %T\n", u1, u1)
	fmt.Printf("u1:%#v, %T\n", u2, u2)
}
