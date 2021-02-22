package repositories

import "goWeb/test/dataModels"

// 需要实现的接口
type MovieRepository interface {
	GetMovieName() string
}

// Movie 的封装类  数据管理员
type MovieManager struct {
}

// 构造函数创建MovieManager, 返回的值的接口 由于MovieManager实现了这个接口
func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	// 模拟已经从查询到数据
	movie := &dataModels.Movice{Name: "理查德姑妈"}
	return movie.Name
}
