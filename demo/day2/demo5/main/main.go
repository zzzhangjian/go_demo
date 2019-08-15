package main

import (
	"fmt"
	"os"
)

func main() {
	egid := os.Getegid()
	fmt.Println("egid:", egid)
	gopath := os.Getenv("GOPATH")
	fmt.Println("GOPATH env:", gopath)

	java_home := os.Getenv("JAVA_HOME")
	fmt.Println("java_home env:", java_home)

}
