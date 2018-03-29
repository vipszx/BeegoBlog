package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"log"
)

type Comment struct {
	Id       int64
	Article  *Article `orm:"rel(fk)"`
	UserName string
	Content  string   `orm:"size(1000)"`
	Created  time.Time
}

func (this *Comment) CommentCount() (count int64) {
	o := orm.NewOrm()
	o.Using("default")

	count, err := o.QueryTable(this).Count()
	if err != nil {
		log.Println("获取评论总数出错：" + err.Error())
	}
	return count
}

func (this *Comment) Create(query *Comment) error {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Insert(query)
	if err != nil {
		return err
	}
	return nil
}

func (this *Comment) GetAll(article_id int64) (result []Comment) {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Raw("SELECT * FROM `comment` WHERE `article_id` = ?", article_id).QueryRows(&result)

	if err != nil {
		log.Println("查询所有文章出错：" + err.Error())
	}
	return result
}



func (this *Comment) Delete(id int64) error {
	o := orm.NewOrm()
	o.Using("default")

	this.Id = id
	_, err := o.Delete(this)
	if err != nil {
		return err
	}
	return nil
}
