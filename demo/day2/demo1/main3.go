package demo1

import "fmt"

func main(){
	printNum(9)
}

func printNum(num int){
	for i:=0;i<=num;i++{
		fmt.Printf("%d + %d = %d \n",num-i,i,num)
	}
}
