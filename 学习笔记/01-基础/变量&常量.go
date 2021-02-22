package main

import "fmt"

// go 中区分大小写，不可重复声明

// 定义变量
var a int

// 定义并赋值
var b int = 233

// 类型推断
var c = "abc"

// 定义变量组
var (
	d int
	e string = "1234"
	f        = 2.3
)

// =====================================

// 定义常量，与变量定义方式一样， 定义是必须赋值
const A int = 123
const B = 1.1
const (
	C string = "abc"
)

// 常量枚举, 在下一次遇到 const 关键字 iota 重置
const (
	D = iota //0
	E        //1
	F        //2
	G        //3
)

func main() {
	fmt.Println(a, b, c, d, e, f)

	fmt.Println(A, B, C)

	fmt.Println("====================")

	fmt.Println(D, E, F, G)

	// 变量 简短定义	:=
	aa := 123
	fmt.Println(aa)

	// 匿名变量 _
	// _ 不可读，不可访问，可重复声明
	bb, _ := 1, 2
	_, cc := 3, 4
	fmt.Println(bb, cc)
}
