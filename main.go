package main

import (
	"github.com/astaxie/beego"

	_ "github.com/willstudy/sniffing/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetLogger("file", `{"filename": "logs/sniffing.log"}`)
	beego.SetLogFuncCall(true)
	beego.Run()
}
