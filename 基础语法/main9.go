// 在协程中使用信道
package main

import (
	"fmt"
)

// var wait2 = sync.WaitGroup{}
// var lock = sync.Mutex{}

func Func() {

}

func main() {
	// 声明一个string信道
	// var ch chan interface{} = make(chan interface{}, 2) // 一次最多只能放2个
	var ch chan string = make(chan string, 2) // 一次最多只能放2个
	ch <- "凤凤"                                // 写入数据到信道中
	ch <- "知道"

	s, ok := <-ch // 从信道读取数据
	fmt.Println(s, ok)
	// fmt.Println(strings.Split(s.(string), ","), ok) 如果是空接口，调用字符串的方法Split，需要.(string)把它断言为字符串

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	close(ch)
}
