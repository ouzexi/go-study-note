// 泛型切片
package main

import "fmt"

type mySlice[T int | string] []T

func main() {
	v1 := mySlice[int]{1, 2, 3}
	v2 := mySlice[string]{"c", "x", "k"}

	fmt.Printf("%T, %v", v1[0], v1[0])
	fmt.Printf("%T, %v", v2[0], v2[0])
}
