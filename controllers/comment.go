package controllers

import (
	"log"
	"BeegoBlog/models"
	"time"
	"strconv"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Store() {
	var query models.Comment
	query.UserName = c.GetString("user_name")
	query.Content = c.GetString("content")
	query.Created = time.Now()
	article_id, _ := c.GetInt64("article_id")
	query.Article = &models.Article{
		Id: article_id,
	}

	err := query.Create(&query)
	if err != nil {
		log.Println("达标评论出错：" + err.Error())
	}
	id := strconv.Itoa(int(article_id))
	c.Ctx.Redirect(302, "/article/"+id)
}

func (c *CommentController) Delete() {
	id, err := c.GetInt64(":id")
	if err != nil {
		log.Println("获取ID失败：" + err.Error())
	}
	//查出评论所属

	var query models.Comment
	err = query.Delete(id)
	if err != nil {
		log.Println("删除评论出错:" + err.Error())
	}
	//article_id := strconv.Itoa(int(query.Article.Id))
	c.Ctx.Redirect(302, c.Ctx.Request.Referer())
}
