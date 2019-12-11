package core

import (
	config2 "gin-demo/core/config"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type MysqlConnectPool struct {
}

var (
	instance *MysqlConnectPool
	db       *dbx.DB
	once     sync.Once
	err      error
)

func DbInstance() *MysqlConnectPool {
	once.Do(func() {
		instance = &MysqlConnectPool{}
	})
	return instance
}

func (p *MysqlConnectPool) NewDB() {
	dbConfig := config2.GetConfig().Database
	db, err = dbx.Open("mysql", dbConfig.User+":"+dbConfig.Password+"@tcp("+dbConfig.Host+":"+dbConfig.Port+")/"+dbConfig.Db)
	//db, err = dbx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/h5_game")
	if err != nil {
		panic(err.Error())
	}
}

func (p *MysqlConnectPool) GetDB() *dbx.DB {
	return db
}
