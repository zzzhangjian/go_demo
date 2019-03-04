package lib_response

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
)

func Json(r *ghttp.Request,err int,msg string,data ...interface{})  {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.Write(g.Map{
		"err": err,
		"msg": msg,
		"data": responseData,
	})
	r.Exit()
}
