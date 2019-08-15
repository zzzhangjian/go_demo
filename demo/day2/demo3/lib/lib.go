package lib

import "fmt"

var Name string
var Age int

func init() {
	fmt.Println("......init start")
	Age = 18
	Name = "zhangjian"
	fmt.Println("Name = ", Name)
	fmt.Println("Age = ", Age)
	fmt.Println("......init end")
}
