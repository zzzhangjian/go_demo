package main

import "fmt"

func main() {
	var sum int = 100
	println("sum=", sum)
	var c chan int = make(chan int, 1)
	println("c=", c)
	person := Person{
		Name: "zhang",
		Age:  19,
	}
	fmt.Println(person.Name)
	modify(person)
	fmt.Println(person.Name)
	var value int = 99
	modifyPoint(&value)
	fmt.Println(value)
}

func modify(person Person) Person {
	person.Name = "li si"
	person.Age = 29
	return person
}

func modifyPoint(value *int) {
	*value = 100
}

type Person struct {
	Name string
	Age  int
}
