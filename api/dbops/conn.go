package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err error
)

// 初始化数据库连接
func init()  {
	dbConn,err = sql.Open("mysql", "username:pwd@tcp(ip:port)/dbname?charset-utf8")
	if err != nil{
		panic(err.Error())
	}
}
