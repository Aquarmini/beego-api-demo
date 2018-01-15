package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/astaxie/beego"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFilterIndexIndexEmpty(t *testing.T) {
	r, _ := http.NewRequest("GET", "/filter/index", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: 测试首页接口", t, func() {
		Convey("HTTP状态码应为200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("返回数据不能为空", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("返回数据success=false, reason=uid is required", func() {
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(w.Body.String()), &data); err == nil {
				So(data["success"], ShouldBeFalse)
				So(data["reason"], ShouldEqual, "uid is required")
			}
		})
	})
}

func TestFilterIndexIndexNotGraterThanTen(t *testing.T) {
	r, _ := http.NewRequest("GET", "/filter/index?uid=1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: 测试首页接口", t, func() {
		Convey("HTTP状态码应为200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("返回数据不能为空", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("返回数据success=false, reason=uid is required", func() {
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(w.Body.String()), &data); err == nil {
				So(data["success"], ShouldBeFalse)
				So(data["reason"], ShouldEqual, "uid must greater than 10")
			}
		})
	})
}

func TestFilterIndexIndex(t *testing.T) {
	r, _ := http.NewRequest("GET", "/filter/index?uid=11", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: 测试首页接口", t, func() {
		Convey("HTTP状态码应为200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("返回数据不能为空", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("返回数据success=false, reason=uid is required", func() {
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(w.Body.String()), &data); err == nil {
				So(data["success"], ShouldBeTrue)
			}
		})
	})
}
