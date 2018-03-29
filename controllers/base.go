package controllers

import (
	"github.com/astaxie/beego"
	"BeegoBlog/models"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) SideBar()  {
	//定义栏目模型
	var category models.Category
	//栏目总数
	categoryCount := category.CategoryCount()
	this.Data["CategoryCount"] = categoryCount
	//所有栏目
	categories := category.GetAll()
	this.Data["Categories"] = categories

	var article models.Article

	//文章总数
	articleCount := article.ArticleCount()
	this.Data["ArticleCount"] = articleCount
	//最新5篇文章
	newArticles := article.GetNew()
	this.Data["NewArticles"] = newArticles

	//定义评论模型获取评论总数
	var comment models.Comment
	commentCount := comment.CommentCount()
	this.Data["CommentCount"] = commentCount
}


