package main

import (
	"fmt"
	"github.com/gogf/gf/g/database/gdb"
)

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname",gconv:"u_nickname"`
}

func main() {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host:     "127.0.0.1",
				Port:     "3306",
				User:     "root",
				Pass:     "root",
				Name:     "weixin_chat",
				Type:     "mysql",
				Role:     "master",
				Priority: 100,
			},
		},
		"user-center": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host:     "192.168.1.110",
				Port:     "3306",
				User:     "root",
				Pass:     "123456",
				Name:     "test",
				Type:     "mysql",
				Role:     "master",
				Priority: 100,
			},
		},
	})

	db, err := gdb.New("default")
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	result, err := db.Table("user").All()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.ToJson())
	}

	user, err := db.Table("user").Where(User{Nickname: ""}).One()
	if err == nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.ToJson())
	}
}
