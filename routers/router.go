package routers

import (
	"boke/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/addarticle", &controllers.ArticleController{},"get:AddArticle;post:StoreArticle")
	beego.Router("/content", &controllers.ArticleController{},"get:ContentArticle")
	beego.Router("/editarticle", &controllers.ArticleController{},"get:EditArticle;post:StoreEditArticle")
	beego.Router("/deletearticle", &controllers.ArticleController{},"get:DeleteArticle")
}
