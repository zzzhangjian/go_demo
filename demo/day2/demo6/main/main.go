package main

func main() {
	var sum int = 100
	println("sum=", sum)
	var c chan int = make(chan int, 1)
	println("c=", c)
}
