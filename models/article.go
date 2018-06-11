package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"log"
)

type Article struct {
	Id       int64
	Category *Category  `orm:"rel(fk)"`
	UserName string
	Title    string
	Content  string     `orm:"size(1000)"`
	Comment  []*Comment `orm:"reverse(many)"`
	View     int64
	Created  time.Time
}

func (this *Article) ArticleCount() (count int64) {
	o := orm.NewOrm()
	o.Using("default")

	count, err := o.QueryTable(this).Count()
	if err != nil {
		log.Println("获取栏目总数出错：" + err.Error())
	}
	return count
}

func (this *Article) Create(query *Article) error {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Insert(query)
	if err != nil {
		return err
	}
	return nil
}

func (this *Article) GetOne(id int64) (result Article, err error) {
	o := orm.NewOrm()
	o.Using("default")

	result = Article{Id: id}
	err = o.Read(&result)
	if err != nil {
		return Article{}, err
	}
	return result, nil
}

func (this *Article) GetAll() (result []Article) {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(this).OrderBy("-created").All(&result)
	if err != nil {
		log.Println("查询所有文章出错：" + err.Error())
	}
	return result
}

func (this *Article) Update(query *Article) error {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Update(query)
	if err != nil {
		log.Println("修改栏目出错:" + err.Error())
		return err
	}
	return nil
}

func (this *Article) Delete(id int64) error {
	o := orm.NewOrm()
	o.Using("default")

	this.Id = id
	_, err := o.Delete(this)
	if err != nil {
		return err
	}
	return nil
}

func (this *Article) GetNew() (result []Article) {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(this).OrderBy("-created").Limit(5).All(&result)
	if err != nil {
		log.Println("查询最近文章出错：" + err.Error())
	}
	return result
}
