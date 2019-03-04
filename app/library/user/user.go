package lib_user

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"errors"
	"github.com/gogf/gf/g/util/gvalid"
	"fmt"
	"github.com/gogf/gf/g/os/gtime"
)
const (
	USER_SESSION_MARK = "user_info"
)

var (
	table = g.DB().Table("user")
)

func SignUp(data g.MapStrStr) error  {
	rules := []string {
		"passport  @required|length:6,16#账户不能为空|长度应当在:min到:max之间",
		"password2 @required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间",
		"password  @required|length:6,16|same:password2#密码不能为空|长度应当在:min到:max之间|两次输入的密码不相同",
	}
	if e := gvalid.CheckMap(data,rules); e != nil {
		return errors.New(e.String())
	}
	if _,ok := data["nickname"]; !ok {
		data["nickname"] = data["passport"]
	}
	if !CheckPassport(data["passport"]) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在",data["passport"]))
	}
	if !CheckPassport(data["nickname"]) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在",data["nickname"]))
	}
	if _,ok := data["create_time"]; !ok {
		data["create_time"] = gtime.Now().String()
	}
	if _,err := table.Filter().Data(data).Save(); err != nil {
		return err
	}
	return nil
}

func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains(USER_SESSION_MARK)
}

func SignIn(passport, password string, session *ghttp.Session) error {
	record, err := table.Where("passport=? and password=?", passport, password).One()
	if err != nil {
		return err
	}
	if record == nil {
		return errors.New("账号或密码错误")
	}
	session.Set(USER_SESSION_MARK,record)
	return nil
}

func SignOut(session *ghttp.Session)  {
	session.Remove(USER_SESSION_MARK)
}

func CheckPassport(passport string) bool {
	if i, err := table.Where("passport",passport).Count(); err != nil {
		return false
	} else {
		return i == 0
	}
}

func CheckNickName(nickName string) bool {
	if i, err := table.Where("nickName",nickName).Count(); err != nil {
		return false
	} else {
		return i == 0
	}
}