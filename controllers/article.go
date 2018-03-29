package controllers

import (
	"github.com/astaxie/beego"
	"BeegoDemo/models"
	"html/template"
	"time"
	"log"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Create() {

	//所有栏目
	var category models.Category
	categories := category.GetAll()
	c.Data["Categories"] = categories

	c.Data["XsrfData"] = template.HTML(c.XSRFFormHTML())

	c.Layout = "layout.html"
	c.TplName = "create.html"
}

func (c *ArticleController) Store() {
	var query models.Article
	query.Title = c.GetString("title")
	query.UserName = c.GetString("user_name")
	query.Content = c.GetString("content")
	query.Created = time.Now()
	caregory_id, _ := c.GetInt64("category_id")
	query.Category = &models.Category{
		Id: caregory_id,
	}

	err := query.Create(&query)
	if err != nil {
		log.Println("新增文章出错：" + err.Error())
	}
	id := strconv.Itoa(int(query.Id))
	c.Ctx.Redirect(302, "/article/"+id)

}
