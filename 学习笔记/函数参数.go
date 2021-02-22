package main

import (
	"fmt"
)

type fu func(int) int

func main() {

	test(1, fun)

}

func test(i int, f fu) {
	fmt.Println(i)
	a := f(i)
	fmt.Println(a)
}

func fun(i int) int {
	fmt.Println("fun ====", i)
	return i
}
