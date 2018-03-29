package controllers

import (
	"github.com/astaxie/beego"
	"BeegoDemo/models"
	"time"
	"log"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Store() {
	var query models.Category
	query.Title = c.GetString("title")
	query.Time = time.Now()
	err := query.Create(&query)
	if err != nil {
		log.Println("添加栏目出错:" + err.Error())
	}
	c.Ctx.Redirect(302, "/")

}

func (c *CategoryController) Delete() {
	id, err := c.GetInt64(":id")
	if err != nil {
		log.Println("获取ID失败：" + err.Error())
	}
	var query models.Category
	err = query.Delete(id)
	if err != nil {
		log.Println("删除栏目出错:" + err.Error())
	}
	c.Ctx.Redirect(302, "/")
}
