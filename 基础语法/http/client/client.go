package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 创建实例对象
	client := new(http.Client)
	// 构造请求对象
	request, _ := http.NewRequest("GET", "http://localhost:8080/index", nil)
	// 发请求
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
}
