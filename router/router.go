package router

import (
	"github.com/gogf/gf/g"
	"github.com/zzzhangjian/go_demo/app/controller/user"
)

func init()  {
	g.Server().BindObject("/user",new(ctl_user.Controller))
}
