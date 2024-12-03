// 读取文件
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// ioutil.ReadFile("./h.txt")

	// 读所有
	/* by, err := os.ReadFile("./h.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(by)) */

	// 读取指定长度
	// 0777 权限
	/* file, err := os.OpenFile("./h.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	} */
	/* var readStr []byte = make([]byte, 6) // 读取6个字节，中文为3字节
	n, err := file.Read(readStr)         // n为指针的位置
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	fmt.Println(string(readStr)) */

	// 读取片段
	/* var readStr []byte = make([]byte, 3)
	file.Seek(3, 0) // 从0开始偏移3位读取3个字节,换行可能是2字节
	n, err := file.Read(readStr)
	fmt.Println(n, err)
	fmt.Println(string(readStr)) */

	// 缓存读取
	/* file, _ := os.Open("./h.txt")

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(b), err)
	} */

	// 按指定分隔符读取
	file, _ := os.Open("./h.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadSlice(';')
		if err == io.EOF {
			break
		}
		_len := len(line)
		fmt.Println(string(line[0:_len-1]), err)
	}
}
