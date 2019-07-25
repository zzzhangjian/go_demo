package main

import (
	"fmt"
	"github.com/gogf/gf/g/container/garray"
	"github.com/gogf/gf/g/database/gdb"
)

type Account struct {
	id              int
	ownerId         int
	ownerTable      string
	account         string
	accountType     string
	money           string
	moneyCumulative string
	state           string
	createTime      string
	updateTime      string
}
type ChainAccount struct {
	id                          int
	password                    string
	chainAccount                string
	chainPrivateKey             string
	bankCard                    string
	ownerId                     int
	ownerName                   string
	ownerTable                  string
	levelValue                  string
	moneyFrozen                 string
	moneyAmount                 string
	moneyWithdraw               string
	moneyCumulative             string
	contributionValue           string
	contributionValueCumulative string
	state                       string
	uidCreate                   int
	uidUpdate                   int
	createTime                  string
	updateTime                  string
	version                     string
}

var (
	NORMAL = "NORMAL"
	FREEZE = "FREEZE"
	AVL    = "AVL"
	ACL    = "ACL"
)

func main() {
	db, err := gdb.New("old")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbPro, err := gdb.New("new")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	records, _ := db.Table("chain_account").All()
	rows := garray.New(true)
	for index, item := range records {
		fmt.Println(index, item)
		ownerId := item["owner_id"].String()
		ownerTable := item["owner_table"].String()
		rows.Append(map[string]string{
			"owner_id":         ownerId,
			"owner_table":      ownerTable,
			"account":          item["id"].String(),
			"money":            "0",
			"money_cumulative": "0",
			"state":            FREEZE,
			"account_type":     ACL,
		})
		rows.Append(map[string]string{
			"owner_id":         ownerId,
			"owner_table":      ownerTable,
			"account":          item["id"].String(),
			"money":            "0",
			"money_cumulative": "0",
			"state":            FREEZE,
			"account_type":     AVL,
		})
		//dbPro.Table("accounts").Data(map[string]string{
		//	"owner_id":ownerId,
		//	"owner_table":ownerTable,
		//	"account":item["id"].String(),
		//	"money":"0",
		//	"money_cumulative":"0",
		//	"state":FREEZE,
		//	"account_type":ACL,
		//},map[string]string{
		//	"owner_id":ownerId,
		//	"owner_table":ownerTable,
		//	"account":item["id"].String(),
		//	"money":"0",
		//	"money_cumulative":"0",
		//	"state":FREEZE,
		//	"account_type":AVL,
		//}).Save()
	}

	dbPro.BatchInsert("accounts", rows.Slice(), 1000)
	result, err := dbPro.Table("accounts").All()
	fmt.Println(result)
}

func init() {
	gdb.SetConfig(gdb.Config{
		"old": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host:     "rm-uf648iqm71x4490175o.mysql.rds.aliyuncs.com",
				Port:     "3306",
				User:     "consumer",
				Pass:     "consumer@hpb",
				Name:     "onebillion",
				Type:     "mysql",
				Role:     "master",
				Priority: 100,
			},
		},
		"new": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host:     "rm-uf648iqm71x4490175o.mysql.rds.aliyuncs.com",
				Port:     "3306",
				User:     "consumer",
				Pass:     "consumer@hpb",
				Name:     "onebillion_pro",
				Type:     "mysql",
				Role:     "master",
				Priority: 100,
			},
		},
	})
}
