package main

import (
	_ "boke/routers"
	_ "boke/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

