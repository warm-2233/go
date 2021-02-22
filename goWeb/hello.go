package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	//输出html
	// 请求方式: GET
	// 访问地址: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		// ctx.HTML返回一个html页面，
		ctx.HTML("<h1>hello</h1>")
	})
	//输出字符串
	// 类似于 app.Handle("GET", "/ping", [...])
	// 请求方式: GET
	// 请求地址: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		// ctx.WriteString将向请求方返回一个字符串
		ctx.WriteString("pong")
	})
	//输出json
	// 请求方式: GET
	// 请求地址: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		// ctx表示返回的结果，ctx.JSON即为返回一个json字符串
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	app.Run(iris.Addr(":8080")) //8080 监听端口
}
