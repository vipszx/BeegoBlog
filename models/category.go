package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type Category struct {
	Id      int64
	Title   string
	Article []*Article `orm:"reverse(many)"`
	Time    time.Time
}

func (this *Category) CategoryCount() (count int64) {
	o := orm.NewOrm()
	o.Using("default")

	count, err := o.QueryTable(this).Count()
	if err != nil {
		log.Println("获取栏目总数出错：" + err.Error())
	}
	return count
}

func (this *Category) GetOne(id int64) (result Category, err error) {
	o := orm.NewOrm()
	o.Using("default")

	result = Category{Id: id}
	err = o.Read(&result)
	if err != nil {
		return Category{}, err
	}
	return result, nil
}


func (this *Category) GetAll() (result []Category) {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(this).All(&result, "id", "title")
	if err != nil {
		log.Println("查询所有栏目出错：" + err.Error())
	}
	return result
}

func (this *Category) Create(query *Category) error {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Insert(query)
	if err != nil {
		return err
	}
	return nil
}

func (this *Category) Update(query *Category) error {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Update(query)
	if err != nil {
		log.Println("修改栏目出错:" + err.Error())
		return err
	}
	return nil
}

func (this *Category) Delete(id int64) error {
	o := orm.NewOrm()
	o.Using("default")

	this.Id = id
	_, err := o.Delete(this)
	if err != nil {
		return err
	}
	return nil
}
