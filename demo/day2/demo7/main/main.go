package main

import "fmt"

func main() {
	a := 10
	b := 99

	a, b = b, a
	fmt.Printf("a:%d,b:%d", a, b)
	swap(&a, &b)
	fmt.Printf("a:%d,b:%d", a, b)
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
	return
}
