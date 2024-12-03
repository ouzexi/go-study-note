// 二进制文件读写
package main

import (
	"io"
	"os"
)

func main() {
	// Open已存在的文件，OpenFile不存在的文件,ReadFile直接读取文件
	/* file1, err := os.ReadFile("pic.jpg")
	if err != nil {
		fmt.Println(err)
		return
	} */

	// 再打开一个文件写入进去
	/* err = os.WriteFile("xx1.jpg", file1, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("图片下载成功") */

	// 文件复制
	file2, _ := os.Open("pic.jpg")
	write, _ := os.Create("xxx.jpg")
	io.Copy(write, file2)
}
