package controllers

import (
	"BeegoDemo/models"
	"html/template"
	"time"
	"log"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Create() {
	//调用BaseController获取栏目列表，栏目总数、文章总数等右边栏数据
	c.SideBar()
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


func (c *ArticleController) Show(){

	id, err:= c.GetInt64(":id")
	if err != nil {
		log.Println("获取ID失败：" + err.Error())
	}
	var article models.Article
	article ,err = article.GetOne(id)
	if err != nil {
		log.Println("查询文章出错:" + err.Error())
	}

	//浏览量+1
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("UPDATE `article` SET `view` = `view` + 1 WHERE `id` = ?",id).Exec()
	//赋值
	c.Data["Article"] = article

	//定义评论模型获取文章的所有李评论
	var comment models.Comment
	comments := comment.GetAll(id)
	c.Data["Comments"] = comments

	//调用BaseController获取栏目列表，栏目总数、文章总数等右边栏数据
	c.SideBar()

	c.Data["XsrfData"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "show.html"
}


func (c *ArticleController) Delete() {
	id, err := c.GetInt64(":id")
	if err != nil {
		log.Println("获取ID失败：" + err.Error())
	}
	var query models.Article
	err = query.Delete(id)
	if err != nil {
		log.Println("删除文章出错:" + err.Error())
	}
	c.Ctx.Redirect(302, "/")
}


func (c *ArticleController) Edit(){
	//调用BaseController获取栏目列表，栏目总数、文章总数等右边栏数据
	c.SideBar()
	id, err:= c.GetInt64(":id")
	if err != nil {
		log.Println("获取ID失败：" + err.Error())
	}
	var article models.Article
	article ,err = article.GetOne(id)
	if err != nil {
		log.Println("查询文章出错:" + err.Error())
	}
	c.Data["ThisArticle"] = article

	c.Layout = "layout.html"
	c.TplName = "edit.html"
}
