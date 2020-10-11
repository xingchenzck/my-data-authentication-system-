package db_mysql

import (
	"database/sql"
	"github.com/beego-develop"
	_ origin"github.com/go-sql-driver/mysql"
)
var Db *sql.DB

func ConnectDB()  {
	//1.读取conf配置信息
	  config :=beego.AppConfig
	  dbDrive := config.String("db_driver")
	  dbUser := config.String("db_user")
	  dbPassword := config.String("db_paddword")
	  dbIp := config.String("db_ip")
	  dbName := config.String("db_name")
	//2.组织链接数据库的字符串
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db, err := sql.Open(dbDrive,connUrl)
	if err != nil {
		panic("数据库连接错误，请检查配置")
	}
   Db =db
	//3.链接数据库
	//4.读取数据库链接对象，处理链接结果

	
}
