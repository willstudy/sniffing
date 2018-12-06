package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
    "github.com/astaxie/beego/httplib"

	"github.com/willstudy/sniffing/models"
)

// Operations about object
type IndexController struct {
	beego.Controller
}

func (o *IndexController) Get() {
    var res models.TodayThing
    req := httplib.Get("https://gank.io/api/today")
    str, err := req.String()
    if err != nil {
		fmt.Printf("req failed")
        //beego.Fatal(err)
        return
    }
	if err := json.Unmarshal([]byte(str), &res); err != nil {
        fmt.Printf("Unmarshal failed")
    }
    o.Data["json"] = res
	o.ServeJSON()
}
