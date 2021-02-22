package main

import (
	"fmt"
)

func main() {
	fun1(1)
	fun1("sd")
	fmt.Println("===================")
	fun2(233, "a", 3.14, func() {
		fmt.Println("func ")
	})
}

func fun1(x interface{}) {
	fmt.Println(x)
	v, ok := x.(int)
	if ok {
		fmt.Println(ok, v)
	} else {
		fmt.Println(ok, v)
	}
}
func fun2(x ...interface{}) {
	fmt.Println(x)
	fmt.Printf("\t %T %v %#v \n", x, x, x)

	for i, v := range x {
		fmt.Println(i, v)
		if a, ok := v.(int); ok {
			fmt.Println("int", a)
		} else if a, ok := v.(string); ok {
			fmt.Println("str", a)
		} else if a, ok := v.(float64); ok {
			fmt.Println("float64", a)
		} else if a, ok := v.(func()); ok {
			fmt.Println("func", a)
		}
	}

}
