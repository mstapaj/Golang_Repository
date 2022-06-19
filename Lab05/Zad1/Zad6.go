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
	var liczba string
	var x int
	var tabelaWynikow []rekord
	limitGora := 100
	limitDol := 0
	nick := "bot"
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
	licznik := 1
	rand.Seed(time.Now().UnixNano())
	wynik := rand.Intn(101)
	liczba = strconv.Itoa(rand.Intn(101))
	for {
		if x, _ = strconv.Atoi(liczba); x > wynik {
			limitGora, _ = strconv.Atoi(liczba)
			liczba = strconv.Itoa((limitGora + limitDol) / 2)
			//liczba = strconv.Itoa(rand.Intn(limitGora-limitDol+1) + limitDol)
			licznik += 1
		} else if x, _ = strconv.Atoi(liczba); x < wynik {
			limitDol, _ = strconv.Atoi(liczba)
			liczba = strconv.Itoa((limitGora + limitDol) / 2)
			//liczba = strconv.Itoa(rand.Intn(limitGora-limitDol+1) + limitDol)
			licznik += 1
		} else if x, _ = strconv.Atoi(liczba); x == wynik {
			break
		}
	}
	fmt.Println("Zgadłeś liczbę za", licznik, "razem")
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
	fmt.Println("Tabela wyników:")
	for i, v := range tabelaWynikow {
		fmt.Println(i+1, v.nickname, "-", v.proby, "-", v.data)
	}
	f, err = os.Create("scores.txt")
	check(err)
	var temp string
	for _, v := range tabelaWynikow {
		temp = v.nickname + ";" + strconv.Itoa(v.proby) + ";" + v.data + "\n"
		_, err := f.Write([]byte(temp))
		check(err)
	}
}
