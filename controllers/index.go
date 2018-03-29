package controllers

import (
	"BeegoDemo/models"
	"html/template"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	//调用BaseController获取栏目列表，栏目总数、文章总数等右边栏数据
	c.SideBar()
	// 定义文章模型
	var article models.Article
	//所有文章
	articles := article.GetAll()
	c.Data["Articles"] = articles


	c.Data["XsrfData"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "index.html"
}
