package main

import (
	"fmt"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/big"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/exchange?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//exist := db.HasTable("identities")
	//exist := db.HasTable(&Identity{})
	exist := db.HasTable(&Account{})
	fmt.Println("%s is exist:", "identities", exist)

	fmt.Println("connect success")
	defer db.Close()
}

type Identity struct {
	gorm.Model
	Email          string
	PasswordDigest string
	IsActive       bool
	RetryCount     uint
	IsLocked       bool
	LastVerifyAt   gtime.Time
}

type Account struct {
	gorm.Model
	MemberId   uint
	CurrencyId uint
	Balance    big.decimal
	Locked     big.decimal
}
