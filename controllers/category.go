package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kevinxu001/goblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	op := this.Input().Get("op")

	switch op {
	case "add":
		cname := this.Input().Get("cname")
		if len(cname) == 0 {
			break
		}

		err := models.AddCategory(cname)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return
	}

	this.Data["IsCategory"] = true
	this.TplNames = "category.html"
	
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
