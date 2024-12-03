// 约束方法
package main

import (
	"fmt"
	"strconv"
)

type Price int

type Price2 string

type showPrice interface {
	~int | ~string
	String() string
}

func showPriceFunc[T showPrice](p T) {
	fmt.Println(p)
}

func (p Price) String() string {
	// int转数字
	return strconv.Itoa(int(p))
}

func (p Price2) String() string {
	return string(p)
}

func main() {
	var p1 Price = 12
	var p2 Price2 = "56"
	// showPriceFunc[Price](p1)
	fmt.Println(p1.String())
	fmt.Println(p2.String())
	showPriceFunc(p1)
}
