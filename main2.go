package main

import (
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"
)

type ControllerTemplate struct {
	gmvc.Controller
}

func (c *ControllerTemplate) Info() {
	c.View.Assign("name", "jhon")
	c.View.Assigns(map[string]interface{}{
		"age":   11,
		"score": 88.5,
	})
	c.View.DisplayContent(`
		<html>
			<head>
			<title>gf template engine</title>
			</head>
				<body>
					<p>Name: {{.name}}</p>
					<p>Age:  {{.age}}</p>
					<p>Score:{{.score}}</p>
				</body>
		</html>
	`)

}
func (c *ControllerTemplate) Index() {
	c.Response.Write("index.html")

}
func main2() {
	s := ghttp.GetServer()
	s.BindController("/template", new(ControllerTemplate))
	s.SetPort(8080)
	s.Run()
}
