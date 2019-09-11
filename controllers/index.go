package controllers

import (
	"boke/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get()  {
	o := orm.NewOrm()
	var article []models.Article
	_, e := o.QueryTable("article").All(&article)
	if e != nil {
		beego.Info("查询出错")
		return
	}
	c.Data["articles"] = article
	c.TplName = "index.html"
}
