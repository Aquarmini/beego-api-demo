package filter

import (
	"github.com/astaxie/beego"
)

// Operations about object
type IndexController struct {
	beego.Controller
}

func (o *IndexController) Index() {
	result := map[string]interface{}{
		"success": true,
	}
	o.Data["json"] = result
	o.ServeJSON()
}
