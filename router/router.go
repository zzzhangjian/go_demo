package router

import (
	"github.com/gogf/gf/g"
	"../app/controller/user"
)

func init()  {
	g.Server().BindObject("/user",new(ctl_user.Controller))
}
