package main

import (
	_ "beegoApi/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.RunMode = "test"

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
