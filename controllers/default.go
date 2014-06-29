package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsIndex"] = true

	this.TplNames = "index.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)
}
