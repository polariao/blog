package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get()  {
	c.TplName = "index.html"
}

func (c *IndexController) Post()  {

	c.Redirect("/index",302)
}
