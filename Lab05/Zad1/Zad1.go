package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var liczba int
	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem z przedziału od 0 do 100")
	rand.Seed(time.Now().UnixNano())
	wynik := rand.Intn(101)
	for {
		fmt.Println("Podaj liczbę: ")
		fmt.Scan(&liczba)
		if liczba > wynik {
			fmt.Println("Za duża liczba")
		} else if liczba < wynik {
			fmt.Println("Za mała liczba")
		} else if wynik == liczba {
			fmt.Println("Gratulacje!")
			break
		} else {
			fmt.Println("Podana liczba jest błędna")
		}
	}
}
