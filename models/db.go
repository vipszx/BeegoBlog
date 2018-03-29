package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

func init() {
	//注册模型
	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(Article))
	orm.RegisterModel(new(Comment))
	//orm.RegisterModel(new(User))
}

//注册数据库
func RegisterDB() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//从配置文件中获取mysql内容
	dbUser := beego.AppConfig.String("dbUser")
	dbPwd := beego.AppConfig.String("dbPwd")
	dbName := beego.AppConfig.String("dbName")
	dbLink := fmt.Sprintf("%s:%s@/%s?charset=utf8&loc=Local", dbUser, dbPwd,dbName)

	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", dbLink)

	//自动建表
	//orm.RunSyncdb("default", false, true)
}
