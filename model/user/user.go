package user

import (
	"fmt"
	"sync"

	user "github.com/li-zeyuan/my-micro-service/proto/user"
)

var (
	s *service
	m sync.RWMutex
)

type service struct {
}

// 用户服务类
type Service interface {
	// 根据用户名获取用户
	QueryUserByName(userName string) (ret *user.User, err error)
}

// 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("Service 未初始化")
	}

	return s, nil
}

// 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = new(service)
}
