package router

import (
	"github.com/astaxie/beego"
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
