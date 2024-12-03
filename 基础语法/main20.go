// 文件写入
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 如果文件不存在就创建写入
	// os.O_RDWR会覆盖 os.O_APPEND会追加
	file, err := os.OpenFile("w.txt", os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写入
	// n, _ := file.WriteString("疯疯知道")
	// n, _ := file.Write([]byte("疯疯11知道"))
	// fmt.Println(n)

	buf := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		// 把格式数据写成字符串
		buf.WriteString(fmt.Sprintf("%d %s\n", i+1, "知道知道"))
	}
	// 提交
	buf.Flush()

}
