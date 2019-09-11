package routers

import (
	"boke/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/index", &controllers.IndexController{})
}
