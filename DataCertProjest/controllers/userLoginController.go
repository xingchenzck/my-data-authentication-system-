package controllers

import (
	"DataCertProjest/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get()  {
	l.TplName = "login.html"
}
func (l *LoginController)Post()  {
	var user models.User
	err :=l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("用户信息解析错误，请重试")
		return
	}
	u , err :=user.QueryUser()
	if err != nil{
	fmt.Println(err.Error())
	l.Ctx.WriteString("抱歉，用户登入失败")
		return
	}
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"
}