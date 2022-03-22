package main

import (
	"fmt"
	"os"
	"strconv"
)

func showResult(num1, num2 int) {
	fmt.Print("Wynik sumowania: ")
	fmt.Println(num1 + num2)
}

func main() {
	var num1, num2 int
	var add string
	if len(os.Args) == 1 {
		fmt.Println("Czy chcesz teraz podać liczby do zsumowania? t/n")
		fmt.Scan(&add)
		if add == "t" {
			fmt.Print("Podaj pierwszą liczbę: ")
			fmt.Scan(&num1)
			fmt.Print("Podaj drugą liczbę: ")
			fmt.Scan(&num2)
			showResult(num1, num2)
		}
	} else if len(os.Args) == 3 {
		if os.Args[1] == "-liczba1" || os.Args[1] == "-liczba2" {
			num1, _ = strconv.Atoi(os.Args[2])
			fmt.Println("Czy chcesz teraz podać druga liczbe do zsumowania? t/n")
			fmt.Scan(&add)
			if add == "t" {
				fmt.Print("Podaj drugą liczbę: ")
				fmt.Scan(&num2)
				showResult(num1, num2)
			}
		} else {
			fmt.Println("Złe argumenty!")
		}
	} else if len(os.Args) == 5 {
		if os.Args[1] == "-liczba1" {
			num1, _ = strconv.Atoi(os.Args[2])
			if os.Args[3] == "-liczba2" {
				num2, _ = strconv.Atoi(os.Args[4])
				showResult(num1, num2)
			} else {
				fmt.Println("Złe argumenty!")
			}
		} else {
			fmt.Println("Złe argumenty!")
		}
	} else {
		println("Podano złe argumenty")
		println("Prawidłowa składnia: go run main.go -liczba1 <liczba1> -liczba2 <liczba2>")
	}
}
