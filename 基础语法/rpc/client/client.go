package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	Num1 int
	Num2 int
}

type Res struct {
	Num int
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	req := Req{1, 2}
	var res Res
	// Server.Add是rpc服务对象的Add方法,客户端远程调用
	client.Call("Server.Add", req, &res)
	fmt.Print(res)
}
