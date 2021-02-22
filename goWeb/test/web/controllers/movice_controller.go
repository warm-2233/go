package controllers

import (
	"goWeb/test/repositories"
	"goWeb/test/service"

	"github.com/kataras/iris/mvc"
)

type MovieController struct {
}

// 这个controller有Get方法的监听, 最后返回一个视图
func (c *MovieController) Get() mvc.View {
	// 创建一个Movie的管理者,他可以实例出movie也可以取movie的属性,相当于movie的管理员
	movieRepository := repositories.NewMovieManager() // 用来使得 repositories 管理 dataModelss
	// 初始化一个服务管理员, 调用他的接口功能 来操作数据管理员
	movieService := service.NewMovieServiceManger(movieRepository)
	// 其实本质上是 数据管理员去给服务管理员movie的名字
	movieResult := movieService.ShowMoviceName()
	return mvc.View{ // 模板实例
		Name: "movie/index.html",
		Data: movieResult,
	}
}
