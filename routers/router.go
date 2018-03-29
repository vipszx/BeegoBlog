package routers

import (
	"BeegoDemo/controllers"
	"github.com/astaxie/beego"
)

//func init() {
//    beego.Router("/", &controllers.MainController{})
//}


//注册路由
func RegisterRouter(){
	//首页
	beego.Router("/", &controllers.IndexController{},"get:Index")

	//栏目路由
	beego.Router("/category", &controllers.CategoryController{},"post:Store")
	beego.Router("/category/delete/:id", &controllers.CategoryController{},"get:Delete")

	//文章路由
	beego.Router("/article/create", &controllers.ArticleController{},"get:Create")
	beego.Router("/article", &controllers.ArticleController{},"post:Store")
	//beego.Router("/category/:id", &controllers.ArticleController{},"get:Show")
	//beego.Router("/category/edit/:id", &controllers.ArticleController{},"get:Edit")
	//beego.Router("/category/:id", &controllers.ArticleController{},"put:Update")
	//beego.Router("/category/delete/:id", &controllers.ArticleController{},"get:Delete")
}