package filter

import (
	"github.com/astaxie/beego"
	"beego-api-demo/filter/test"
)

func Register() {

	beego.InsertFilter("/filter/*", beego.BeforeRouter, test.UserFilter)
}
