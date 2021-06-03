package mysql

import (
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	client Collections
	logger logs.Collections
}

var mysqlClient MySQL

func GenerateUri(host, port, name, username, password string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, name)
}

func InitConnection(masterDBUrl string, logger logs.Collections) MySQL {
	//openConnDBMaster, err := gorm.Open("mysql", masterDBUrl)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//mysqlClient = MySQL{
	//	client: openConnDBMaster,
	//	logger: logger,
	//}
	//
	//// Globally mode
	//sqlite
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
	//	Logger: newLogger,
	//})
	//
	//logger.Info("MySQL Connected")
	//
	//return mysqlClient
	return MySQL{}
}
