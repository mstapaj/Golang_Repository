package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	type rekord struct {
		proby    int
		nickname string
	}
	var liczba, kontynuuj, nick string
	var x int
	var tabelaWynikow []rekord
	for {
		licznik := 1
		fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem z przedziału od 0 do 100")
		fmt.Println("Napisz koniec, aby zakończyć działanie programu i wyświetlić poprawną odpowiedź")
		rand.Seed(time.Now().UnixNano())
		wynik := rand.Intn(101)
		for {
			fmt.Print("Podaj liczbę: ")
			fmt.Scan(&liczba)
			if liczba == "koniec" {
				fmt.Println("Koniec programu")
				fmt.Println("Wynik to:", wynik)
				break
			} else if x, _ = strconv.Atoi(liczba); x > wynik {
				fmt.Println("Za duża liczba")
				licznik += 1
			} else if x, _ = strconv.Atoi(liczba); x < wynik {
				fmt.Println("Za mała liczba")
				licznik += 1
			} else if x, _ = strconv.Atoi(liczba); x == wynik {
				fmt.Println("Gratulacje!")
				break
			} else {
				fmt.Println("Podana liczba jest błędna")
			}

		}
		fmt.Println("Zgadłeś liczbę za", licznik, "razem")
		fmt.Print("Podaj swój nick:")
		fmt.Scan(&nick)
		tabelaWynikow = append(tabelaWynikow, rekord{nickname: nick, proby: licznik})
		fmt.Println("Czy gramy jeszcze raz? [T/N]")
		fmt.Scan(&kontynuuj)
		if strings.ToUpper(kontynuuj) == "N" {
			sort.Slice(tabelaWynikow, func(p, q int) bool {
				return tabelaWynikow[p].proby < tabelaWynikow[q].proby
			})
			fmt.Println("Tabela wyników:")
			for i, v := range tabelaWynikow {
				fmt.Println(i+1, v.nickname, "-", v.proby)
			}
			break
		}
	}
}
