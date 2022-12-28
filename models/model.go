package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"

	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

// User -
type User struct {
	ID       int    `orm:"column(id)"`
	Name     string `orm:"column(name)"`
	Password string `orm:"column(password)"`
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(User))
	// 注册模型
	orm.RegisterModel(new(Order))
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// need to register default database
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/mydatabase?charset=utf8")
	if err != nil {
		fmt.Println(1111, "MYSQL", err)
	}
}
