package router

import (
	"github.com/astaxie/beego"
	"time"
)

// Operations about object
type IndexController struct {
	beego.Controller
}

func (o *IndexController) Index() {
	result := map[string]string{
		"version": beego.AppConfig.String("version"),
	}
	result["data"] = "You can fly with beego!"
	o.Data["json"] = result
	o.ServeJSON()
}

func (o *IndexController) Time() {
	t := time.Now()
	result := map[string]int64{
		"time": t.Unix(),
	}
	o.Data["json"] = result
	o.ServeJSON()
}
