package user

import (
	"github.com/go-log/log"
	"github.com/li-zeyuan/my-micro-service/basic/db"
	proto "github.com/li-zeyuan/my-micro-service/proto/user"
)

func (s *service) QueryUserByName(userName string) (*proto.User, error) {
	queryString := `select user_id, user_name, pwd from user where user_name = ?`

	// 获取数据库
	o := db.GetDB()

	ret := new(proto.User)

	// 查询
	err := o.QueryRow(queryString, userName).Scan(ret.Id, ret.Name, ret.Pwd)
	if err != nil {
		log.Logf("get user info error")
		return ret, err
	}

	return ret, nil
}
