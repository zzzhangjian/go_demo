package main

import (
	"fmt"
	"time"
)

var (
	Man    = 1
	Female = 2
)

func main() {
	for {
		printTimeNow()
		time.Sleep(1000 * time.Millisecond)
	}
}

func printTimeNow() {
	now := time.Now()
	fmt.Println(now)
	if now.Second()%Female == 0 {
		fmt.Println("now is Female", Female)
	} else {
		fmt.Println("now is Man", Man)
	}
}
