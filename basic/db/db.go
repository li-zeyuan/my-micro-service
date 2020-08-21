package db

import (
	"database/sql"
	"fmt"
	"github.com/go-log/log"
	"github.com/li-zeyuan/my-micro-service/basic/config"
	"sync"
)

var (
	inited bool
	mysqlDB *sql.DB
	m sync.RWMutex
)

// 初始化数据库
func Init()  {
	m.Lock()
	defer m.Unlock()

	if inited {
		err := fmt.Errorf("mysql 已经初始化过了")
		log.Log(err)
		return
	}

	// 如果配置声明使用MySQL
	if config.GetMysqlConfig().GetEnabled(){
		initMysql()
	}

	inited = true
}

// 获取db
func GetDB() *sql.DB {
	return mysqlDB
}