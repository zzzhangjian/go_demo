package demo1

import (
	"fmt"
	"time"
)
var sum  = 0
func main(){
	for i :=0;i<10;i++{
		go print(i)
	}
	time.Sleep(10)
	fmt.Println(sum)
}

func print(num int)  {
	fmt.Println(num)
	sum += num
}
