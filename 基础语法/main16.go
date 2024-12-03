// 泛型约束
package main

import "fmt"

type NumStr interface {
	Num | Str
}

type Num interface {
	~int | ~int32 | ~int64 | ~uint8
}

type Str interface {
	string
}

type Status uint8

type myySlice[T NumStr] []T

func main() {
	m1 := myySlice[int]{1, 2, 3}
	m2 := myySlice[string]{"1", "2", "3"}
	m3 := myySlice[int32]{11, 22, 33}
	// m4 := myySlice[Status]{11, 22, 33}
	// 如果uint8不用波浪线，那么以uint8作为别名（底层数据类型）的类型Status是不被允许的

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
}
