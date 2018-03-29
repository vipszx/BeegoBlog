package controllers

import (
	"github.com/astaxie/beego"
	"BeegoDemo/models"
	"html/template"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {

	//定义栏目模型
	var category models.Category
	//栏目总数
	categoryCount := category.CategoryCount()
	c.Data["CategoryCount"] = categoryCount

	//所有栏目
	categories := category.GetAll()
	c.Data["Categories"] = categories

	// 定义文章模型
	var article models.Article
	//所有文章
	articles := article.GetAll()

	c.Data["Articles"] = articles

	c.Data["ArticleCount"] = 20

	c.Data["XsrfData"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "index.html"

}
