package controllers

import (
	"boke/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get()  {
	c.TplName = "login.html"
}

func (c *LoginController) Post()  {
	username := c.GetString("username")
	password := c.GetString("password")
	o := orm.NewOrm()
	user := new(models.User)
	user.Name = username
	err := o.Read(user,"name")
	if err != nil{
		c.Data["code"] = "用户不存在!"
		c.TplName = "login.html"
		return
	}
	if password != user.Password {
		c.Data["code"] = "密码不正确!"
		c.TplName = "login.html"
		return
	}
	c.Redirect("/index",302)
}
