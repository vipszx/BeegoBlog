package models

import "time"

type Comment struct {
	Id       int64
	Article  *Article `orm:"rel(fk)"`
	UserName string
	Content  string   `orm:"size(1000)"`
	Created  time.Time
}

