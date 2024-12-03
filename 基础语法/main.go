package main

import (
	"fmt"
	. "go_study/pkg"
	"go_study/version"
)

const (
	age int = 1
	cnt int = 2
	def int = iota
)

func deferUtil() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次调用接收到n=%v\n", n)
		i++
		fmt.Printf("匿名工具函数第%v次被调用\n", i)

		fmt.Print(version.Version)
		// fmt.Print(version.pkgName)
		return i
	}
}

type UserInfo struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" bingding:"min=3,max=6"`
}

/* func Defer() int {
	f := deferUtil()
	//f(1)
	defer f(1)
	//defer f(f(3))
	return f(2)
  } */

func Defer() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

func testChan() {
	// 声明一个string信道，容量为2
	var ch chan string = make(chan string, 2)

	ch <- "枫枫" // 写入数据到信道中
	ch <- "知道"

	s := <-ch // 从信道读取数据
	fmt.Println(s)
	ss, ok := <-ch
	fmt.Println(ss, ok)
	close(ch)
}

func pushNum(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) // 写完必须要关闭，不然会死锁
}

type User struct {
	Name string
	Age  int
}

func Print(inter interface{}) {
	switch x := inter.(type) {
	case User:
		x.Name = "张三"
		fmt.Println(x.Name, x.Age)
	}
}

func PPrint(user *User) {
	user.Name = "王五"
}

func main() {
	// Defer()
	// pkg.Pkg()
	// var arr [3]int = [3]int{7, 8, 9}
	Pkg()
	/* fmt.Println("xxx")
	var arr = [5]int{4, 5, 6, 7, 8}
	var slice = arr[0:4] // 切片的两边也是顾头不顾尾
	slice[0] = 9         // 如果改变切片中的元素，对应的数组也会变化
	fmt.Println(arr)

	var s []int
	s = make([]int, 3, 5)
	fmt.Println(cap(s)) */

	/* s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s1 = make([]int, 8, 8)
	copy(s1, s) // 将s复制到s1，复制的个数就是s1的容量大小
	fmt.Println(s1)

	str := "hello 世界"
	fmt.Printf("[]byte(str)=%s\n", []byte(str))
	fmt.Printf("[]byte(str)=%v\n", []byte(str))
	// 汉字被分成了三个数字
	for i, v := range str {
		fmt.Printf("str[%d]=%c\n", i, v)
	} */

	// testChan()

	/* var c1 chan int = make(chan int, 2) // 2表示缓冲区大小
	go pushNum(c1)

	for value := range c1 {
		fmt.Println(value)
	} */

	user := User{"枫枫", 21}
	Print(user)
	fmt.Println(&user)
}
