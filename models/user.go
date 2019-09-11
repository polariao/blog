package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	Name string `orm:"unique"`
	Password string
}

type Article struct {
	ArtId  int `orm:"pk;auto"`
	ArtName string `orm:"null;unique"`
	ArtTime time.Time `orm:"auto_now;type(datetime)"`
	ArtCount int `orm:"default(0)"`
	ArtContent string `orm:"null"`
	ArtImg string `orm:"null"`
}


func init()  {
	//设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:100521@/first?charset=utf8")
	//注册定义的model
	orm.RegisterModel(new(User),new(Article))
	//生成表
	orm.RunSyncdb("default",false,true)    //force为true,将回滚数据表结构，同时清空之前数据，一般设为;false
}

