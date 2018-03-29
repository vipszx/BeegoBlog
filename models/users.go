package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"log"
)

type User struct {
	Id      int64
	Name    string     `orm:"unique"`
	Email   string     `orm:"unique"`
	Created time.Time
	Article []*Article `orm:"reverse(many)"`
	Comment []*Comment `orm:"reverse(many)"`
}

func init() {
	//orm.RegisterModel(new(User))
}

func (this *User) GetUsers() (result []User) {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(this).All(&result)
	if err != nil {
		log.Println("查询用户出错：" + err.Error())
	}
	return
}

func (this *User) Insert(query *User) error {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Insert(query)
	if err != nil {
		log.Println("添加用户出错:" + err.Error())
		return err
	}
	return nil
}

func (this *User) Update(query *User) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Update(query)
	if err != nil {
		log.Println("修改用户出错:" + err.Error())
		return err
	}
	return nil
}

func (this *User) Delete(id int64) error {
	o := orm.NewOrm()
	o.Using("default")

    this.Id = id
	_, err := o.Delete(this)
	if err != nil {
		log.Println("删除用户出错:" + err.Error())
		return err
	}
	return nil
}
