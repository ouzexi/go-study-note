// 泛型map
package main

import "fmt"

type myMap[K string | int, V any] map[K]V

type Userr struct {
	Name string
}

func main() {
	m1 := myMap[string, string]{
		"key": "ff",
	}
	m2 := myMap[int, Userr]{
		0: Userr{
			Name: "haha",
		},
	}
	fmt.Println(m1)
	fmt.Println(m2)
}
