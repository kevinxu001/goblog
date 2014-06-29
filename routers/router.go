package routers

import (
	"github.com/astaxie/beego"
	"github.com/kevinxu001/goblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}
