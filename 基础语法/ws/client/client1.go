package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	di := websocket.Dialer{}
	conn, _, err := di.Dial("ws://127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 传递 类型 和 内容
	conn.WriteMessage(websocket.TextMessage, []byte("你好啊"))

	// 传送消息
	go Send(conn)

	for {
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(msgType, string(msgData))
	}

	// 断开
	conn.Close()
}

func Send(conn *websocket.Conn) {
	for {
		// 接收命令行参数
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}
}
