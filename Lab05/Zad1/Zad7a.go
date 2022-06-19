package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var response string
	limitGora := 100
	limitDol := 0
	rand.Seed(time.Now().UnixNano())
	liczba := strconv.Itoa(rand.Intn(101))
	response, err := reader.ReadString('\n')
	check(err)
	response, err = reader.ReadString('\n')
	check(err)

	fmt.Println(liczba)
	for {
		response, err := reader.ReadString('\n')
		check(err)
		if response == "Podaj liczbe:\n" {
			fmt.Println(liczba)
		}
		if response == "Za duża liczba\n" {
			limitGora, _ = strconv.Atoi(liczba)
			liczba = strconv.Itoa((limitGora + limitDol) / 2)
		} else if response == "Za mała liczba\n" {
			limitDol, _ = strconv.Atoi(liczba)
			liczba = strconv.Itoa((limitGora + limitDol) / 2)
		} else if response == "Gratulacje!\n" {
			break
		}
	}
	response, err = reader.ReadString('\n')
	check(err)
	response, err = reader.ReadString('\n')
	check(err)
	if response == "Podaj swój nick:\n" {
		fmt.Println("bot")
	}

	response, err = reader.ReadString('\n')
	check(err)
	if response != "Czy gramy jeszcze raz? [T/N]\n" {
		response, err = reader.ReadString('\n')
		check(err)
		if response != "Czy gramy jeszcze raz? [T/N]\n" {
			response, err = reader.ReadString('\n')
			check(err)
		}
	}
	if response == "Czy gramy jeszcze raz? [T/N]\n" {
		fmt.Println("N")
	}
}
