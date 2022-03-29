package main

import (
	"fmt"
)

func prime(a int) bool {
	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func primeRange(a int, b int) {
	for i := a; i < b; i++ {
		if prime(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	var a, b int
	fmt.Print("Wprowadź liczbę startową: ")
	fmt.Scan(&a)
	fmt.Print("Wprowadź liczbę końcową: ")
	fmt.Scan(&b)
	fmt.Println()
	primeRange(a, b)
}
