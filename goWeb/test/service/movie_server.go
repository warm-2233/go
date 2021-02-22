package service

import (
	"fmt"
	"goWeb/test/repositories"
)

// 服务提供的接口,用来整合repositories, 服务管理员的功能接口
type MovieServer interface {
	ShowMoviceName() string
}

// 服务管理员, 用来管理数据管理员
type MovieServiceManager struct {
	repo repositories.MovieRepository
}

// 初始化服务管理员
func NewMovieServiceManger(repo repositories.MovieRepository) MovieServer {
	return &MovieServiceManager{repo: repo}
}

// 服务管理员的功能实现
func (m *MovieServiceManager) ShowMoviceName() string {
	fmt.Println("我们获取到的视频名称： " + m.repo.GetMovieName())
	return "我们获取到的视频名称： " + m.repo.GetMovieName()
}
