package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kevinxu001/goblog/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplNames = "topic.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Topics"], err = models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")

	tid := this.Input().Get("tid")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, content)
	} else {
		err = models.ModifyTopic(tid, title, category, content)
	}
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	this.Data["IsTopic"] = true
	this.TplNames = "topic_add.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) View() {
	this.Data["IsTopic"] = true
	this.TplNames = "topic_view.html"

	topic, err := models.GetTopic(this.Ctx.Input.Params["0"])
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
}

func (this *TopicController) Modify() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.Data["IsTopic"] = true
	this.TplNames = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Ctx.Input.Params["0"])
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/", 302)
}
