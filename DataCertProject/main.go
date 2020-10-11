package main

import (
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"

)

func main() {
	db_mysql.ConnectDB()
    //加载静态文件
    beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")


	beego.Run()
}

