package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type player struct {
	chanel chan string
	name   string
}

func pingPong(ch1 chan string, ch2 chan string, wg *sync.WaitGroup, msg string) {
	for {
		<-ch1
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
		ch2 <- "Ball"
		wg.Done()
	}
}

func makeTabPlayers(amount int) []player {
	tab := make([]player, 0)
	for i := 0; i < amount; i++ {
		temp := player{name: "PingPong " + strconv.Itoa(i), chanel: make(chan string)}
		tab = append(tab, temp)
	}
	return tab
}

func playPingPong(tabOfPlayers []player, iteration int) {
	var wg sync.WaitGroup
	wg.Add(iteration)
	for i, player := range tabOfPlayers {
		if i+1 > len(tabOfPlayers)-1 {
			go pingPong(player.chanel, tabOfPlayers[0].chanel, &wg, player.name)
		} else {
			go pingPong(player.chanel, tabOfPlayers[i+1].chanel, &wg, player.name)
		}

	}
	tabOfPlayers[0].chanel <- "Ball"
	wg.Wait()
}

func main() {
	tab := makeTabPlayers(20)
	playPingPong(tab, 100)
}
