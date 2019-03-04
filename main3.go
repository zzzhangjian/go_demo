package main

import (
	_ "./router"
	_ "./boot"
	"github.com/gogf/gf/g"
)
func main(){
	g.Server().Run()
}