// 在协程中使用信道
package main

import (
	"fmt"
	"sync"
)

var ch chan int = make(chan int, 10)
var wait3 = sync.WaitGroup{}

func insertNum() {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	wait3.Done()
}

func main() {

	// 声明可读channel
	var read <-chan int = ch
	// 声明可写channel
	var write chan<- int = ch

	write <- 1
	write <- 2

	fmt.Println(<-read)
	fmt.Println(<-read)

	wait3.Add(2)

	go insertNum()
	go insertNum()

	wait3.Wait()
	close(ch)
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(i, ok)
	}
}
