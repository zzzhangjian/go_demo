package boot

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

func init() {
	v := g.View()
	c := g.Config()
	s := g.Server()

	c.AddPath("config")
	v.AddPath("template")
	v.SetDelimiters("${", "}")

	logpath := c.GetString("setting.logpath")
	glog.SetPath(logpath)
	glog.SetStdPrint(true)

	s.SetServerRoot("public")
	s.SetLogPath(logpath)
	s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_ALLLOWER)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	s.SetPort(8080)

}
