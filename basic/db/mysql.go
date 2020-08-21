package db

import (
	"database/sql"
	"github.com/li-zeyuan/my-micro-service/basic/config"
	"log"
)

func initMysql()  {
	// 创建连接
	mysqlDB, err := sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	// 激活链接
	if err := mysqlDB.Ping(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}