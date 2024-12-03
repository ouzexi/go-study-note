package version

import "fmt"

func main() {
	var age int
	// var arr [3]int = [3]int{1, 2, 3};
	// var arr2 = [...]int{2,3,4,5}
	// arr3 := [...]int{1,3,4,5,6}
	fmt.Println("请输入你的年龄：")
	fmt.Scan(&age)
	if age < 18 {
		fmt.Print("未成年")
	} else {
		fmt.Print("已成年")
	}
}
