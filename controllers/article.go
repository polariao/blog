package controllers

import (
	"boke/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"path"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) AddArticle()  {
	c.TplName = "add.html"
}

func (c *ArticleController) StoreArticle()  {
	artname := c.GetString("artname")
	artcontent := c.GetString("artcontent")
	file, header, e := c.GetFile("artfile")
	if artname == "" ||artcontent== "" {
		c.Data["code"] = "数据不完整!"
		c.TplName = "add.html"
		return
	}
	defer file.Close()
	if e != nil {
		c.Data["code"] = "文件上传失败！"
		c.TplName = "add.html"
		return
	}
	//1、限制格式
	ext := path.Ext(header.Filename)   //获取后缀名
	if ext != ".jpg" && ext != ".png" && ext != ".gif" {
		c.Data["code"] = "文件格式不正确！"
		c.TplName = "add.html"
		return
	}
	//2、限制大小
	if header.Size > 50000000{
		c.Data["code"] = "文件过大！"
		c.TplName = "add.html"
		return
	}
	//3、重命名，防止同文件名，后者覆盖前者
	unix := strconv.FormatInt(time.Now().Unix(),10) + ext
	beego.Info(unix)

	if e != nil {
		beego.Info("上传失败",e)
	}else{
		//进行存储
		c.SaveToFile("artfile","./static/img/"+unix)
	}
	//    文件上传结束
	url := "/static/img/"+ unix
	o := orm.NewOrm()
	article := new(models.Article)
	article.ArtName = artname
	article.ArtContent = artcontent
	article.ArtImg = url
	_, err := o.Insert(article)
	if err!= nil {
		c.Data["code"] = "添加文章失败！"
		c.TplName = "add.html"
		return
	}
	c.Redirect("/index",302)
}

func (c *ArticleController) ContentArticle()  {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	article := new(models.Article)
	article.ArtId = id
	o.Read(article,"art_id")
	c.Data["article"] = article
	c.TplName = "content.html"
}

func (c *ArticleController) EditArticle()  {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	article := new(models.Article)
	article.ArtId = id
	o.Read(article,"art_id")
	c.Data["article"] = article
	c.TplName = "edit.html"
}

func (c * ArticleController) StoreEditArticle() {
	var url string
	id, _ := c.GetInt("id")
	artname := c.GetString("artname")
	artcontent := c.GetString("artcontent")
	//文件部分
	file, header, e := c.GetFile("artfile")
	defer file.Close()
	//判断是否上传文件
	if header != nil {
		if e != nil {
			c.Data["code"] = "文件上传失败！"
			c.TplName = "add.html"
			return
		}
		//1、限制格式
		ext := path.Ext(header.Filename) //获取后缀名
		if ext != ".jpg" && ext != ".png" && ext != ".gif" {
			c.Data["code"] = "文件格式不正确！"
			c.TplName = "add.html"
			return
		}
		//2、限制大小
		if header.Size > 50000000 {
			c.Data["code"] = "文件过大！"
			c.TplName = "add.html"
			return
		}
		//3、重命名，防止同文件名，后者覆盖前者
		unix := strconv.FormatInt(time.Now().Unix(), 10) + ext
		beego.Info(unix)

		if e != nil {
			beego.Info("上传失败", e)
		} else {
			//进行存储
			c.SaveToFile("artfile", "./static/img/"+unix)
			//    文件上传结束
			url = "/static/img/" + unix
		}
	}
	o := orm.NewOrm()
	article := new(models.Article)
	article.ArtId = id
	err := o.Read(article, "art_id")
	if err != nil {
		c.Data["code"] = "文数据错误！"
		c.TplName = "index.html"
		return
	}else {
		if article.ArtName != artname && artname != "" {
			article.ArtName = artname
		}
		if article.ArtContent != artcontent && artcontent != "" {
			article.ArtContent = artcontent
		}
		if url!="" {
			article.ArtImg = url
		}
		_, err := o.Insert(article)
		if err!=nil {
			c.Data["code"] = "更新失败！"
			c.TplName = "index.html"
			return
		}
		c.Redirect("/index",302)
	}
}
