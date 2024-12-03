package main

import (
	"fmt"
	"net"
)

func main() {
	// 建立连接
	conn, _ := net.DialTCP("tcp", nil, &net.TCPAddr{net.ParseIP("127.0.0.1"), 8080, ""})

	// 向服务端发送请求
	// conn.Write([]byte("申请出战"))
	// conn.Close()

	// 可以通过命令行输入

	for {
		var s string
		fmt.Scanln(&s)
		if s == "q" {
			break
		}
		conn.Write([]byte(s))

		// 接收服务端传来的值
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println("接收服务端传来的值:" + string(buf[:n]))
	}
	conn.Close()
}
