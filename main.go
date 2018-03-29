package main

import (
	_ "BeegoDemo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"BeegoDemo/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"BeegoDemo/routers"
)

func init() {
	//注册数据库
	models.RegisterDB()
}

func main() {
	log.SetFlags(log.Llongfile)

	// 开启 ORM 调试模式
	orm.Debug = true

	//var user models.User
	//allUser := user.GetUsers()
	//log.Println(allUser)
	//
	//user.Id = 1
	//user.Name = "levitate111"
	//user.Email = "123@14.com"
	//user.Created = time.Now()
	//
	//err := user.Delete(user.Id)
	//if err != nil {
	//	log.Println("删除用户出错：" + err.Error())
	//}


	//注册路由
	routers.RegisterRouter()

	beego.Run()
}
