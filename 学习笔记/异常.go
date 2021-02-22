package main

import (
	"fmt"
)

func main() {

	fmt.Println("start")
	c()
	// b()
	fmt.Println("end")
}

func a() {
	fmt.Println("a")

	panic("怎么写的代码")
}

func b() {
	fmt.Println("b")
	defer func() {
		e := recover()
		fmt.Println(e)
	}()
	a()
}

func c() {
	fmt.Println("c")
	b()
}
