package main

import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//var response string
	//limitGora := 100
	//limitDol := 0
	//rand.Seed(time.Now().UnixNano())
	//liczba := strconv.Itoa(rand.Intn(101))
	//for {
	//	fmt.Print(liczba)
	//	_, err := fmt.Scan(&response)
	//	check(err)
	//	if response == "Za duża liczba\n" {
	//		limitGora, _ = strconv.Atoi(liczba)
	//		liczba = strconv.Itoa((limitGora + limitDol) / 2)
	//		//liczba = strconv.Itoa(rand.Intn(limitGora-limitDol+1) + limitDol)
	//	} else if response == "Za mała liczba\n" {
	//		limitDol, _ = strconv.Atoi(liczba)
	//		liczba = strconv.Itoa((limitGora + limitDol) / 2)
	//		//liczba = strconv.Itoa(rand.Intn(limitGora-limitDol+1) + limitDol)
	//	} else if response == "Gratulacje!\n" {
	//		break
	//	}
	//}
	//for {
	//	_, err := fmt.Scan(&response)
	//	check(err)
	//	if response == "Podaj swój nick:" {
	//		fmt.Print("bot")
	//	}
	//	if response == "Czy gramy jeszcze raz? [T/N]" {
	//		fmt.Print("N")
	//		break
	//	}
	//}
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
