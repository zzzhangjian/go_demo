package main

import "fmt"

func main() {

	for i := 101; i <= 200; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}

}

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
