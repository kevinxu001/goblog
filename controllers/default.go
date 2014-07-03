package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kevinxu001/goblog/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsIndex"] = true

	this.TplNames = "index.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Topics"], err = models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
}
