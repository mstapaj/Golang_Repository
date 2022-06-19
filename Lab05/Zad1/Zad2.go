package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var liczba string
	var x int
	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem z przedziału od 0 do 100")
	fmt.Println("Napisz koniec, aby zakończyć działanie programu i wyświetlić poprawną odpowiedź")
	rand.Seed(time.Now().UnixNano())
	wynik := rand.Intn(101)
	for {
		fmt.Println("Podaj liczbę: ")
		fmt.Scan(&liczba)
		if liczba == "koniec" {
			fmt.Println("Koniec programu")
			fmt.Println("Wynik to:", wynik)
			break
		} else if x, _ = strconv.Atoi(liczba); x > wynik {
			fmt.Println("Za duża liczba")
		} else if x, _ = strconv.Atoi(liczba); x < wynik {
			fmt.Println("Za mała liczba")
		} else if x, _ = strconv.Atoi(liczba); x == wynik {
			fmt.Println("Gratulacje!")
			break
		} else {
			fmt.Println("Podana liczba jest błędna")
		}
	}
}
