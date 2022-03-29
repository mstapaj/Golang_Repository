package main

import "fmt"

func prime(a int) bool {
	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func goldbach(a int) {
	if a%2 != 0 {
		fmt.Println("Podana liczba musi byc liczba parzysta")
	} else {
		for i := 3; i < a; i++ {
			if prime(i) && prime(a-i) {
				fmt.Printf("%1d = %1d + %1d \n", a, i, a-i)
				break
			}
		}
	}
}

func main() {
	goldbach(28)
}
