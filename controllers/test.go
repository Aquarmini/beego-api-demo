package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about object
type TestController struct {
	beego.Controller
}

func (o *TestController) Hello() {
	result := map[string]string{
		"version": "v1.0.0",
	}
	result["data"] = "You can fly with beego!"
	o.Data["json"] = result
	o.ServeJSON()
}
