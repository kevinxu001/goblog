package controllers

import (
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplNames = "topic.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Add() {
	this.Data["IsTopic"] = true
	this.TplNames = "topic_add.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)
}
