package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// 把http升级为websocket
var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connList []*websocket.Conn

func handler(res http.ResponseWriter, req *http.Request) {
	// 服务升级
	conn, err := UP.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 连接建立后把用户append进入
	connList = append(connList, conn)
	for {
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		// 群发给在线用户
		for _, c := range connList {
			// 数据元丰不动返回
			c.WriteMessage(msgType, []byte(fmt.Sprintf("你说的是：%s", msgData)))
		}
		// fmt.Println(msgType, string(msgData))
	}
	defer conn.Close()
	fmt.Println("服务关闭")
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
