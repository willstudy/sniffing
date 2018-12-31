package main

import (
	"database/sql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
    _ "github.com/go-sql-driver/mysql"

	_ "github.com/willstudy/sniffing/routers"
	"github.com/willstudy/sniffing/models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/sniffing.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	logs.SetLogFuncCall(true)
	// init mysql
	var err error
	dsn := beego.AppConfig.String("mysqldsn")
	models.Sdb, err = sql.Open("mysql", dsn)
	if err != nil {
		logs.Error("init mysql failed, error:", err)
		return
	}
	err = models.Sdb.Ping()
    if (err != nil) {
        models.Sdb.Close()
		logs.Error("ping mysql failed")
		return
	}
	logs.Debug("init mysql success")
	beego.Run()
}
