package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	_ "beego-api-demo/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
)

func TestRouterIndexIndex(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: 测试首页接口", t, func() {
		Convey("HTTP状态码应为200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("返回数据不能为空", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("返回版本号为1.0.0", func() {
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(w.Body.String()), &data); err == nil {
				So(data["version"], ShouldEqual, "1.0.0")
			}
		})
	})
}

func TestRouterIndexTime(t *testing.T) {
	r, _ := http.NewRequest("GET", "/router/time", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: 测试时间戳", t, func() {
		Convey("HTTP状态码应为200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("返回数据不能为空", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("返回时间戳大于0", func() {
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(w.Body.String()), &data); err == nil {
				So(data["time"], ShouldBeGreaterThan, 0)
			}
		})
	})
}
