package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("rand int is %d \n", rand.Int())
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("rand int is %f \n", rand.Float32())
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("rand int is %d \n", rand.Intn(100))
	}

}
