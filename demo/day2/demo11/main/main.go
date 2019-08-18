package main

import (
	"fmt"
	"github.com/gogf/gf/g/text/gstr"
)

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var str string = "hello"
	var str2 string = "world"
	var s = str + str2
	fmt.Println(str + str2)
	fmt.Println(s)

	s = fmt.Sprintf("%s %s", str, str2)
	fmt.Println(s)

	substr := s[0:7]
	fmt.Println(substr)

	fmt.Println(s[:5])

	fmt.Println(gstr.Reverse(substr))
	fmt.Println(reverse(s))
}
