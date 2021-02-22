package main

import (
	"fmt"
)

func main() {
	var i *int
	var a int
	fmt.Println(i, a)
	fmt.Println(&a)
	i = &a
	fmt.Println(i)
	fmt.Println(*i)

	fmt.Println("====================================")

	var b = &Test{1, 2}

	b.abc()
	fmt.Println(b.a, b.b)
	b.cba()
	fmt.Println(b.a, b.b)
}

type Test struct {
	a, b int
}

func (t Test) abc() {
	t.a = 111
	t.b = 222
	fmt.Println(t.a, t.b)
}
func (t *Test) cba() {
	t.a = 1111
	t.b = 2222
	fmt.Println(t.a, t.b)
}
