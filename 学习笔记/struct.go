package main

import "fmt"

func main() {
	var t = &a{233, "abcd"}
	fmt.Println(t)
	t.add()
	fmt.Println(t)

	fmt.Println("========================")
	tt := &b{&a{233, "abcd"}, 1234}
	fmt.Println(tt)
	tt.add()

	fmt.Println(tt.a)
}

type a struct {
	x int
	y string
}

type b struct {
	*a
	c int
}

func (aa *a) add() {
	aa.x += 1
	fmt.Println(aa)
}

func (b *b) add() {
	b.a.add()
	b.x += 100
	fmt.Println("b")
}
