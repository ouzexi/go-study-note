package main

import (
	"fmt"
	"sync"
	"time"
)

// 可以通过切片或者map获取协程的返回值
var wait1 = sync.WaitGroup{}
var list []string
var mmap map[string]interface{} = map[string]interface{}{}

func GetName(funName string) {
	time.Sleep(1 * time.Second)
	list = append(list, "返回值")
	mmap[funName] = "返回值 " + funName
	wait1.Done()
}

func main() {
	fmt.Println(list)
	wait1.Add(2)
	go GetName("nameFunc")
	go GetName("newFunc")
	wait1.Wait()
	fmt.Println(list)
	fmt.Println(mmap["nameFunc"])
	fmt.Println(mmap["newFunc"])
}
