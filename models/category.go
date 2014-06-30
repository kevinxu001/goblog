package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

func AddCategory(cname string) error {
	o := orm.NewOrm()

	cate := &Category{Title: cname}

	qs := o.QueryTable("category")
	err := qs.Filter("title", cname).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}
