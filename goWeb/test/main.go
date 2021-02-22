package main

import (
	"log"

	"goWeb/test/web/controllers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// 模板注册  将符合路径下的以指定结尾的文件作为模板
	app.RegisterView(iris.HTML("./web/views", ".html"))
	// 注册控制器, 请求过来后触发对应的controller
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))

	err := app.Run(
		iris.Addr("localhost:8080"),
	)
	if err != nil {
		log.Fatal(err)
	}
}
