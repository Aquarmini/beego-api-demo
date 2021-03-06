// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego-api-demo/controllers"
	"beego-api-demo/controllers/router"
	"github.com/astaxie/beego"
	filterController "beego-api-demo/controllers/filter"
	"beego-api-demo/filter"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	filter.Register()

	// 路由基本使用，返回版本号
	beego.Router("/", &router.IndexController{}, "*:Index")
	// 返回当前时间戳
	beego.Router("/router/time", &router.IndexController{}, "*:Time")
	// 测试ParseForm方法
	beego.Router("/router/parse/form", &router.IndexController{}, "*:ParseFormAction")
	// 测试Filter过滤器
	beego.Router("/filter/index", &filterController.IndexController{}, "*:Index")
}
