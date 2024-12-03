package main

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}
var wait = sync.WaitGroup{}
var num = 0

func AddNum() {
	// 上一把锁，谁先抢到，谁就往下执行
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num++
	}
	// 执行完成之后，解开锁，其他的协程就会去抢
	lock.Unlock()
	wait.Done()
}

func main() {
	wait.Add(2)

	// 因为2个协程执行的时间相近，cpu可能会把多个请求合并为一次，导致执行结果不对
	go AddNum()
	go AddNum()

	wait.Wait()
	fmt.Println(num)
}
