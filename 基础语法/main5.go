package main

import "fmt"

// 空接口没有任何方法，所以相当于所有类型都实现了它，类似TS的any类型(除了切片)
type empty interface{}

type Animal interface {
	Sing()
	Dance(time int)
	Rap() string
}

type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Println("Chicken 调用了Sing 方法")
}

func (c Chicken) Dance(time int) {
	fmt.Println("Chicken 调用了Dance 方法")
}

func (c Chicken) Rap() string {
	fmt.Println("Chicken 调用了Rap 方法")
	return "rap"
}

func Sing(animal Animal) {
	animal.Sing()
}

func main() {
	/* var animal Animal
	animal = Chicken{
		Name: "cxk",
	}
	animal.Sing() */

	// 多态
	chicken := Chicken{Name: "ikun"}
	// cat := Cat{Name: "ali"}
	Sing(chicken)
	// Sing(cat)
}
