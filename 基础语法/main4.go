package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Article struct {
		Title    string `json:"title"` // 序列化后传给前端的字段为title
		Content  string `json:"-"`     // -也不参与序列化
		Author   string
		password string //小写开头的字段不会序列化
	}

	article := Article{
		Title:    "哈哈哈",
		Content:  "内容内容内容",
		Author:   "ouzx",
		password: "1111",
	}

	jsonData, err := json.Marshal(article)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}
