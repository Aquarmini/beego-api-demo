package test

import (
	"github.com/astaxie/beego/context"
	"strconv"
)

var UserFilter = func(ctx *context.Context) {
	uid := ctx.Request.FormValue("uid")
	result := map[string]interface{}{}

	if uid == "" {
		// 未传入UID
		result["success"] = false
		result["reason"] = "uid is required"
		ctx.Output.JSON(result, false, false)
	}

	uid32, _ := strconv.Atoi(uid)

	if uid32 < 10 {
		// uid 小于 10
		result["success"] = false
		result["reason"] = "uid must greater than 10"
		ctx.Output.JSON(result, false, false)
	}
}
