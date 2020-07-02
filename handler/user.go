package handler

import (
	"context"
	"github.com/li-zeyuan/my-micro-service/model/user"
	"log"
)

var (
	userService user.Service
)

type Service struct {

}

// 初始化handler
func Init()  {
	userService, err := user.GetService()
	if err != nil {
		log.Fatal("get user service error")
		return
	}
}

func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error{
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = new(s.Error)
		rsp.Code = 500
		rsp.Detail = err.Error()

		return err
	}

	rsp.User = user
	rsp.Success = true

	return nil
}