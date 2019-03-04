package main

import (
	_ "github.com/zzzhangjian/go_demo/router"
	_ "github.com/zzzhangjian/go_demo/boot"
	"github.com/gogf/gf/g"
)
func main(){
	g.Server().Run()
}