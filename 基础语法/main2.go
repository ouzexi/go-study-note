package main

import (
	"fmt"
)

type msgType uint16

func main() {
	/* var m map[string][]string = map[string][]string{
		"name":  []string{"c", "x", "k"},
		"likes": []string{"唱", "跳", "rap"},
	} */

	/* m["name"][2] = "狗"
	fmt.Println(m["name"])
	fmt.Println(m) */

	var successMsg msgType = 1000
	var errorMsg msgType = 1210

	var num uint16 = 23

	// successMsg = msgType(num)

	fmt.Printf("%v, %T\n", successMsg, successMsg)
	fmt.Printf("%v, %T\n", errorMsg, errorMsg)
	fmt.Printf("%v, %T\n", num, num)
}
