package main

import (
	"fmt"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/gtime"
	"time"
)

type User struct {
	Id               int    `gconv:"id"`
	LoginId          string `gconv:"u_login_id"`
	Nickname         string `gconv:"u_nick_name"`
	Password         string `gconv:"u_password"`
	FriendshipPolicy int    `gconv:"u_friendship_policy_id"`
	State            int    `gconv:"u_user_state_id"`
	PolicyQuestion   string `gconv:"u_friend_policy_question"`
	PolicyAnswer     string `gconv:"u_friend_policy_answer"`
	PolicyPassword   string `gconv:"u_friend_policy_password"`
	RegTime          string `gconv:"u_reg_time"`
}

type UserFriendshipPolicy struct {
	Id               int    `gconv:"id"`
	FriendshipPolicy string `gconv:"u_friendship_policy"`
}

type UserState struct {
	Id   int    `gconv:"id"`
	Name string `gconv:"us_name"`
}

func main() {
	db, err := gdb.New()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var ufp = UserFriendshipPolicy{
		FriendshipPolicy: "NEED_CHECK",
	}
	recordUserFriendshipPolicy, err := db.Table("user_friendship_policy").Data(ufp).Save()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(recordUserFriendshipPolicy)

	var state = UserState{
		Name: "NORMAL",
	}

	recordState, err := db.Table("user_state").Data(state).Save()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(recordState)

	//var cstZone = time.FixedZone("CST", 8*3600)
	local, _ := time.LoadLocation("Asia/Shanghai")
	regTime := time.Now().In(local)
	fmt.Println(regTime)

	var user = User{
		LoginId:          "wx_132",
		Nickname:         "nickname",
		Password:         "password",
		FriendshipPolicy: ufp.Id,
		State:            state.Id,
		RegTime:          gtime.Datetime(),
	}
	fmt.Println(user)
	db.Table("user").Data(user).Save()

	result, err := db.Table("user").Where("u_login_id", "wx_132").One()
	if err != nil {
		fmt.Println(err)
	} else {
		for i, item := range result {
			fmt.Println(i, item)
		}

	}

	u, err := db.Table("user").Where(User{Nickname: "jjjjjj"}).One()
	if err == nil {
		fmt.Println(err)
	} else {
		fmt.Println(u.ToJson())
	}
}

func init() {
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
}
