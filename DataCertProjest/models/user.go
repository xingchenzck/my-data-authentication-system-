package models

import (
	"DataCertProjest/db_mysql"
	"DataCertProjest/util"
)

type User struct {
	Id    int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}

func (u User)SaveUser() (int64 ,error)  {
	u.Password = util.MD5HashString(u.Password)

	row ,err := db_mysql.Db.Exec("insert into user(phone, password)"+
		" values(?,?) ", u.Phone, u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (u User) QueryUser()(*User,error)  {
     u.Password = util.MD5HashString(u.Password)

     row := db_mysql.Db.QueryRow("select phone from user where  phone = ? and password = ?",
		 u.Phone, u.Password)
     err :=row.Scan(&u.Phone)
	if err != nil {
		return &u,nil
	}
return &u,nil

}
