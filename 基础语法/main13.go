package main

import "fmt"

// 用[T type1 | type2 | ...]声明传参的类型
func PrintSlice[T int | int64 | string](slice []T) {
	for _, v := range slice {
		fmt.Printf("%T %v\n", v, v)
	}
}

func main() {
	PrintSlice([]int{1, 2, 3, 4, 5})
	var int64Slice []int64 = []int64{4, 5, 7}
	PrintSlice(int64Slice)
	PrintSlice([]string{"c", "x", "k"})
}
