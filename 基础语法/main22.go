// 目录操作
package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.ReadDir("../go-study")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(dir)
	for _, enter := range dir {
		fmt.Println(enter.Name())
		fmt.Println(enter.Type()) // 权限
		fmt.Println(enter.IsDir())
		// enter.Info()
	}
}
