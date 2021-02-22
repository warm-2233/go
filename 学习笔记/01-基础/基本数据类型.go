package main

import "fmt"

// 字符串
var s string = "abc"
var s2 = `原始字符串` // 原始字符串，不接受任何转意

// 整数型
var i int = 123 // int8 int16 int32 int64    对应位数的整形， int 对于操作系统的位数

// 无符号整型
var u uint = 123 // uint8 uint16 uint32 uint64    对应位数的无符号整型， uint 对于操作系统的位数

// 浮点型
var f float32 = 3.14 // float32， float64  对于位数精度的浮点数

// 布尔型
var t bool = true // 布尔型，只有 true/false 两个值

// 类型别名
var r rune = 'a' // rune == int32  'a'==97
var b byte = 123 // byte == int8
// 自定义
type MyInt = int // 自定义类型别名

// complex64   complex128    复数类型

func main() {
	fmt.Println(s, s2)
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(u)
	fmt.Println(t)
	fmt.Println("===============")
	fmt.Println(r)
	fmt.Println(b)

	fmt.Println("----------------")
	var a MyInt = 233
	fmt.Println(a)
}
