package handler

import (
	"context"
	"log"

	"github.com/li-zeyuan/my-micro-service/model/user"
	proto_user "github.com/li-zeyuan/my-micro-service/proto/user"
)

var (
	userService user.Service
)

type Service struct {

}

// 初始化handler
func Init()  {
	var err error
	userService, err = user.GetService()
	if err != nil {
		log.Fatal("get user service error")
		return
	}
}

func (e *Service) QueryUserByName(ctx context.Context, req *proto_user.Request, rsp *proto_user.Response) error{
	u, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = new(proto_user.Error)
		rsp.Error.Code = 500
		rsp.Error.Detail = err.Error()

		return err
	}

	rsp.User = u
	rsp.Success = true

	return nil
}