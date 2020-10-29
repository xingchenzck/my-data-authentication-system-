package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func  ConnectDB()  {

	confing :=beego.AppConfig
	dbDriver :=confing.String("db_driver")
	dbUser :=confing.String("db_user")
	dbPassword := confing.String("db_password")
	dbIp :=confing.String("db_ip")
	dbName :=confing.String("db_name")

	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	//3、链接数据库
	db, err := sql.Open(dbDriver,connUrl)
	if err !=nil {
		panic("数据库链接错误，请检查配置")
	}
	Db =db
}
