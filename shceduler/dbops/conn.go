package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var (
	dbConn *sql.DB
	err error
)

// 初始化数据库连接
func init()  {

	dbConn,err = sql.Open("mysql", "root:9830@tcp(localhost:3306)/video_server?charset-utf8")
	if err != nil{
		panic(err.Error())
	}
}
