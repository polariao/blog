package controllers

import (
	"boke/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get()  {
	c.TplName = "register.html"
}

func (c *RegisterController) Post()  {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		c.Data["code"] = "数据不完整！"
		c.TplName = "register.html"
		return
	}
	o := orm.NewOrm()
	user := new(models.User)
	user.Name = username
	user.Password = password
	_, e := o.Insert(user)
	if e != nil {
		c.Data["code"] = "注册失败!"
		c.TplName = "register.html"
		return
	}
	c.Redirect("/login",302)
}
