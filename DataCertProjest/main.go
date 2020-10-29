package main

import (
	"DataCertProjest/blockchain"
	"DataCertProjest/db_mysql"
	_ "DataCertProjest/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	bc :=blockchain.NewBlockChain()
	fmt.Printf("创世区块的Hash值 ：%x\n",bc.LashHash)
	block ,err :=bc.SeveDate([]byte("这里存储上链的数据信息"))
	if err !=nil {
     fmt.Printf(err.Error())
		return
	}

	fmt.Printf("区块的高度：%d\n",block.Height)
	fmt.Printf("区块的PrevHash : %x\n",block.PrevHash)

	//1、链接数据库
	db_mysql.ConnectDB()

	//2、静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

