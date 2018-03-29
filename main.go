package main

import (
	_ "BeegoBlog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"BeegoBlog/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"BeegoBlog/routers"
)

func init() {
	//注册数据库
	models.RegisterDB()
}

func main() {
	log.SetFlags(log.Llongfile)

	// 开启 ORM 调试模式
	orm.Debug = false

	//注册路由
	routers.RegisterRouter()

	beego.Run()
}
