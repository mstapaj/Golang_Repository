package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	adjectives := map[int]string{
		1:  "Łosoś",
		2:  "Dziekan",
		3:  "Hiszpańska inkwizycja",
		4:  "Miodożer",
		5:  "Pies",
		6:  "Kot",
		7:  "Małpa",
		8:  "Gibon",
		9:  "Grosik",
		10: "Baron",
		11: "Baleron",
		12: "Knur",
		13: "Koza",
		14: "Truskawka",
		15: "Opuncja",
		16: "Pomelątko",
		17: "Sigma",
		18: "Kubuś Puchatek",
		19: "Dzwon",
		20: "Jajko",
		21: "Siwy",
		22: "Łysy",
		23: "Kononowicz",
		24: "Tarantyzm",
		25: "Imponderabilia",
		26: "Rudy",
		27: "Złodziej",
		28: "Morderca",
		29: "Delfin",
		30: "Tulipan",
		31: "Arka Gdynia",
	}

	noun := map[rune]string{
		'A': "Miły",
		'B': "Zły",
		'C': "Niebieski",
		'D': "Czerwony",
		'E': "Żółty",
		'F': "Zielony",
		'G': "Fioletowy",
		'H': "Różowy",
		'I': "Pomarańczowy",
		'J': "Szary",
		'K': "Czarny",
		'L': "Biały",
		'Ł': "Głupi",
		'M': "Mądry",
		'N': "Inteligentny",
		'O': "Zazdrosny",
		'P': "Dziwny",
		'R': "Kolorowy",
		'S': "Wolny",
		'T': "Szybki",
		'U': "Mocny",
		'W': "Morski",
		'X': "Powolny",
		'Y': "Krzywy",
		'Z': "Prosty",
		'Ź': "Wąski",
		'Ż': "Gruby",
	}

	verb := map[rune]string{
		'A': "Skacze",
		'B': "Biega",
		'C': "Leży",
		'D': "Gra",
		'E': "Z Afryki",
		'F': "Z Azji",
		'G': "Z Europy",
		'H': "Z Ameryki Południowej",
		'I': "Z Ameryki Północnej",
		'J': "Z Australii",
		'K': "Z Antarktydy",
		'L': "Kopie",
		'Ł': "Pływa",
		'M': "Lata",
		'N': "Pije",
		'O': "Uczy się",
		'P': "Pisze",
		'R': "Czyta",
		'S': "Śpiewa",
		'T': "Koloruje",
		'U': "Myje się",
		'W': "Koduje",
		'X': "Rysuje",
		'Y': "Maluje",
		'Z': "Gotuje",
		'Ź': "Szkicuje",
		'Ż': "Rzeżbi",
	}

	var date int
	var first, second string
	if len(os.Args) == 4 {
		date, _ = strconv.Atoi(os.Args[1])
		first = os.Args[2]
		second = os.Args[3]
	} else if len(os.Args) == 3 {
		date, _ = strconv.Atoi(os.Args[1])
		first = os.Args[2]
		fmt.Print("Podaj nazwisko: ")
		fmt.Scan(&second)
	} else if len(os.Args) == 2 {
		date, _ = strconv.Atoi(os.Args[1])
		fmt.Print("Podaj imię: ")
		fmt.Scan(&first)
		fmt.Print("Podaj nazwisko: ")
		fmt.Scan(&second)
	} else {
		fmt.Print("Podaj dzień urodzenia: ")
		fmt.Scan(&date)
		fmt.Print("Podaj imię: ")
		fmt.Scan(&first)
		fmt.Print("Podaj nazwisko: ")
		fmt.Scan(&second)
	}
	first = strings.ToUpper(first)
	second = strings.ToUpper(second)
	if date > 0 && date < 32 {
		fmt.Print("Pseudonim: ")
		fmt.Println(noun[rune(first[0])], adjectives[date], verb[rune(second[0])])
	} else {
		fmt.Println("Data jest nieprawidłowa")
	}

}
