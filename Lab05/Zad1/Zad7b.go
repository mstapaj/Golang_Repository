package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type rekord struct {
	proby    int
	nickname string
	data     string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func filterSlice(slice []rekord, nickname string) []rekord {
	var tab []rekord
	for _, v := range slice {
		if v.nickname == nickname {
			tab = append(tab, v)
		}
	}
	return tab
}

func sortScores(tabelaWynikow []rekord) {
	sort.Slice(tabelaWynikow, func(p, q int) bool {
		return tabelaWynikow[p].proby < tabelaWynikow[q].proby
	})
}

func main() {
	var liczba, kontynuuj, nick string
	var x int
	var tabelaWynikow []rekord
	f, err := os.Open("scores.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var temp rekord
		temp.nickname = strings.Split(scanner.Text(), ";")[0]
		temp.proby, _ = strconv.Atoi(strings.Split(scanner.Text(), ";")[1])
		temp.data = strings.Split(strings.Split(scanner.Text(), ";")[2], ".")[0]
		tabelaWynikow = append(tabelaWynikow, temp)
	}
	sortScores(tabelaWynikow)
	for {
		licznik := 1
		fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem z przedziału od 0 do 100")
		rand.Seed(time.Now().UnixNano())
		wynik := rand.Intn(101)
		for {
			fmt.Println("Podaj liczbe:")
			_, err := fmt.Scan(&liczba)
			check(err)
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
		fmt.Println("Podaj swój nick:")
		_, err := fmt.Scan(&nick)
		currentTime := time.Now()
		check(err)
		if len(tabelaWynikow) > 0 {
			if tabelaWynikow[0].proby > licznik {
				fmt.Println("Nowy rekord globalny")
			}
		}
		toPB := filterSlice(tabelaWynikow, nick)
		if len(toPB) > 0 {
			if toPB[0].proby > licznik {
				fmt.Println("Nowy rekord personalny")
			}
		} else {
			fmt.Println("Nowy rekord personalny")
		}
		tabelaWynikow = append(tabelaWynikow, rekord{nickname: nick, proby: licznik, data: strings.Split(currentTime.String(), ".")[0]})
		sortScores(tabelaWynikow)
		fmt.Println("Czy gramy jeszcze raz? [T/N]")
		_, err = fmt.Scan(&kontynuuj)
		check(err)
		if strings.ToUpper(kontynuuj) == "N" {
			fmt.Println("Tabela wyników:")
			for i, v := range tabelaWynikow {
				fmt.Println(i+1, v.nickname, "-", v.proby, "-", v.data)
			}
			f, err := os.Create("scores.txt")
			check(err)
			var temp string
			for _, v := range tabelaWynikow {
				temp = v.nickname + ";" + strconv.Itoa(v.proby) + ";" + v.data + "\n"
				_, err := f.Write([]byte(temp))
				check(err)
			}
			break
		}
	}
}
