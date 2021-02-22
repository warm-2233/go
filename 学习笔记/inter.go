package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human  // 匿名字段
	school string
}

// Human 实现 SayHi 方法
func (h Human) SayHi() {
	fmt.Printf("SayHi %s %v\n", h.name, h.age)
}

// Human 实现 Sing 方法
func (h Human) Sing(lyrics string) {
	fmt.Println("======Sing", lyrics)
}

// Interface Men 被 Human, Student 实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	stu := Student{Human{"A", 100}, "bibiibibi"}
	stu.SayHi()
	stu.Sing("=====================")

	var m Men
	m = stu

	m.SayHi()
	m.Sing("mmmmmmmmmmmm")

	var i interface{}
	i = stu
	fmt.Println(i)
}
