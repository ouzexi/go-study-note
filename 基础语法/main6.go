package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 发送验证🐎函数
func SendCode() {
	fmt.Println("发送验证码开始...")
	time.Sleep((3 * time.Second))
	fmt.Println("发送验证码结束...")
	wg.Done() // 协程结束了，或者wg.Add(-1)
}

func main() {
	wg.Add(1) // 分一个协程去执行
	// 实现用户注册功能
	fmt.Println("用户注册校验完成")
	// 发送验证码 (如果不开启协程，需要等待3秒后才打印 注册验证部分结束)，类似同步 / 异步
	// 会阻塞主线程
	// SendCode()
	go SendCode()
	fmt.Println("注册验证部分结束")
	// time.Sleep((4 * time.Second)) // 需要等待协程结束，主线程才结束，否则主线程结束，协程也会跟着结束

	wg.Wait() // 等待协程结束再结束
	fmt.Println("main结束...")
}
