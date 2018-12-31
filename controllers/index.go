package controllers

import (
	"strings"
	"time"
	/*
	"encoding/json"
	"fmt"

    "github.com/astaxie/beego/httplib" */
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/willstudy/sniffing/models"
)

// Operations about object
type IndexController struct {
	beego.Controller
}

func (o *IndexController) Get() {
    var res models.IndexInfo
	res.Code = 0
	res.Msg = "success"

	// index data
	sql := "SELECT img, target_url, thing_id from index_data"
	rows, err := models.Sdb.Query(sql)
	if err != nil {
		logs.Warn("get index data failed, sql:[%s], err:[%v]", sql, err)
		return
	}

	var img string
	var targetUrl string
	var thingID string
	for rows.Next() {
		err = rows.Scan(&img, &targetUrl, &thingID)
		if err != nil {
			logs.Warn("scan row failed, err[%v]", err)
			continue
		}
	}
	// banner
	imgs := strings.Split(img, ",")
	urls := strings.Split(targetUrl, ",")
	for i, _ := range urls {
		var banner models.BannerInfo
		banner.Img = imgs[i]
		banner.Target = urls[i]
		res.Data.Banner = append(res.Data.Banner, banner)
	}

	sql = "SELECT title, create_time, breif, images, publish_time, origin_url,"
	sql += "author,avatar,source,type,commmets,stars from thing where _id in " + thingID
	rows, err = models.Sdb.Query(sql)
	if err != nil {
		logs.Warn("get index data failed, sql:[%s], err:[%v]", sql, err)
		return
	}
	for rows.Next() {
		var hot models.HotInfo
		var c_t, p_t int64
		err = rows.Scan(&hot.Title, &c_t, &hot.Brief, &hot.Image, &p_t, &hot.TargetUrl,
				&hot.Author, &hot.Avatar, &hot.Source, &hot.Type, &hot.Comments, &hot.Stars)
		if err != nil {
				logs.Warn("scan row failed, err[%v]", err)
				continue
		}
		imgArr := strings.Split(hot.Image, ",")
		hot.Image = imgArr[0]
		hot.CreatedAt = time.Unix(c_t, 0).String()
		hot.PublishAt = time.Unix(p_t, 0).String()
		res.Data.Content.Hot = append(res.Data.Content.Hot, hot)
	}
	o.Data["json"] = res
	o.ServeJSON()
}
