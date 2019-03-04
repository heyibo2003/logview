package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
//初始数据链接参数
type Mysql struct {
	mysql_user     string
	mysql_passwd   string
	mysql_database string
	mysql_host     string
	mysql_port     string
}
//数据库链接
func init() {
	var c Mysql
	enable_MySQL, err := beego.AppConfig.Bool("enable_mysql")
	if err == nil && enable_MySQL == true {
		c.mysql_user = beego.AppConfig.String("mysql_user")
		c.mysql_passwd = beego.AppConfig.String("mysql_passwd")
		c.mysql_database = beego.AppConfig.String("mysql_database")
		c.mysql_host = beego.AppConfig.String("mysql_host")
		c.mysql_port = beego.AppConfig.String("mysql_port")
	}
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql",
		c.mysql_user+":"+c.mysql_passwd+"@tcp("+c.mysql_host+":"+c.mysql_port+")/"+c.mysql_database+"?charset=utf8")

}
