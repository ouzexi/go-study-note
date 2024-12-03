package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	// fmt.Println(req)
	data, err := os.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
	}
	// 设置响应头
	header := res.Header()
	header["token"] = []string{"sdsdsdsds"}

	// 获取post请求的body
	// data, err = io.ReadAll(req.Body)

	// var mapa map[string]string
	ma := make(map[string]string)
	json.Unmarshal(data, &ma)
	// res.Write([]byte("hello 枫枫"))
	res.Write(data)
}

func main() {
	// 绑定回调
	http.HandleFunc("/index", IndexHandler)
	// 注册服务
	fmt.Println("listen server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
