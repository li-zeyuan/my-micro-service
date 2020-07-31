package basic

import (
	"github.com/li-zeyuan/my-micro-service/basic/config"
	"github.com/li-zeyuan/my-micro-service/basic/db"
)

func Init()  {
	config.Init()
	db.Init()
}
