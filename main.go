package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kevinxu001/goblog/models"
	_ "github.com/kevinxu001/goblog/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	//orm.RunCommand()
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		beego.Error(err)
	}

	beego.Run()
}
