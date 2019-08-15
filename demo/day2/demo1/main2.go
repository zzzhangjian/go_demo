package demo1

import (
	"fmt"
	"time"
	"github.com/derekparker/delve/service/api"
)

func main(){
	var pipe chan int
	pipe = make(chan int,1)
	go Add(100,300,pipe)
	sum := <- pipe
	fmt.Println(sum)
}

func add(start int,end int,pipe chan int)  {
	for num := start;num<end;num++{

	}
	return
}
