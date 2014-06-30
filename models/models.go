package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DB_DRIVER  = "mysql"
	_DB_HOST    = "10.16.128.9"
	_DB_PORT    = "3306"
	_DB_NAME    = "goblog"
	_DB_USER    = "root"
	_DB_PWD     = "nohacker"
	_DB_CHARSET = "utf8" //root:root@(127.0.0.1:3389)/orm?charset=utf8
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index;auto_now_add"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index;auto_now"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index;auto_now_add"`
	Updated         time.Time `orm:"index;auto_now"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index;null"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_DB_DRIVER, orm.DR_MySQL)
	orm.RegisterDataBase("default", _DB_DRIVER, _DB_USER+":"+_DB_PWD+"@("+_DB_HOST+":"+_DB_PORT+")/"+_DB_NAME+"?charset="+_DB_CHARSET, 10)
}
