package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建TCP监听
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("127.0.0.1"), 8080, ""})
	if err != nil {
		fmt.Println(err)
		return
	}
	// 等待连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String())
		// 接收客户端传来的值
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		fmt.Println("接收客户端传来的值:" + string(buf[:n]))
		// 元丰不动传回客户端吧
		conn.Write(buf[:n])
	}

	// 获取数据，发送数据
}
