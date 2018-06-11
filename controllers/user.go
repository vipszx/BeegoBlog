package controllers

import (
	"github.com/astaxie/beego"
	"BeegoBlog/models"
	"html/template"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Index() {
	c.Data["XsrfData"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "login.html"
	beego.ReadFromRequest(&c.Controller)

}

func (c *UserController) Login() {
	var User models.User
	username := c.GetString("user_name")
	password := c.GetString("password")
	result := User.GetUserByUserName(username)
	if result.Password == password {
		c.SetSession("uid", int(result.Id))
		c.Ctx.Redirect(302, "/")
	}
	flash := beego.NewFlash()
	flash.Error("账号或密码错误")
	flash.Store(&c.Controller)
	c.Ctx.Redirect(302, "/user/login")
}

func (c *UserController) Register() {

}
