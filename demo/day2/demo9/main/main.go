package main

import "fmt"

func main() {
	var str string = "hello word"
	fmt.Print(str)

	var str1 string = `
	窗前明月光，
	低头思故乡。
	举头望明月，
	我是郭德纲。
	`
	fmt.Println(str1)

	var b byte = 'c'
	fmt.Println(b)
	fmt.Printf("%c\n", b)

}
