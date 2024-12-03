package main

import "fmt"

// select case 搭配 for循环 确定何时退出
func main() {
	var ch chan int = make(chan int, 2)
	ch <- 1
	ch <- 2

	// exit:
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		default:
			fmt.Println("没有数据了")
			// break exit
			goto exit
		}
	}
exit:
	close(ch)
}
