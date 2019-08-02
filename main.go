package main

import (
	"github.com/gogf/gf/g/frame/gins"
	"github.com/gogf/gf/g/net/ghttp"
)

func main1() {
	s := ghttp.GetServer()
	s.BindHandler("/template", func(r *ghttp.Request) {
		content, _ := gins.View().Parse("index.html", map[string]interface{}{
			"id":   123,
			"name": "jhon",
			"age":  18,
		})
		r.Response.Write(content)
	})
	s.SetPort(8080)
	s.Run()
}
