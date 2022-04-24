package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type ant struct {
	x       int
	y       int
	hasLeaf bool
	leaf    leaf
}

type leaf struct {
	x      int
	y      int
	hasAnt bool
}

func (a *ant) move(x int, y int) {
	a.x += x
	a.y += y
	if a.hasLeaf {
		a.leaf.x = a.x
		a.leaf.y = a.y
	}
}

func (a *ant) pickLeaf(l leaf) {
	a.hasLeaf = true
	a.leaf = l
	l.hasAnt = true
}

func (l *leaf) putAwayLeaf(x, y int) {
	l.x = x
	l.y = y
	l.hasAnt = false
}

func (a *ant) deleteLeaf(l leaf, x int, y int) {
	a.hasLeaf = false
	l.hasAnt = false
	l.putAwayLeaf(x, y)
}

func (l *leaf) addAnt() {
	l.hasAnt = true
}

func findByCords(x, y int, tab [10]leaf) leaf {
	var res leaf
	for _, l := range tab {
		if l.x == x && l.y == y {
			res = l
		}
	}
	return res
}

func main() {
	tab := [10][10]string{}
	for i, strings := range tab {
		for i2, _ := range strings {
			tab[i][i2] = "_"
		}
	}

	tabAnts := [5]ant{}
	tabLeafs := [10]leaf{}

	rand.Seed(time.Now().UnixNano())

	for i, _ := range tabAnts {
		for {
			x := rand.Intn(10)
			y := rand.Intn(10)
			if tab[x][y] == "_" {
				tabAnts[i] = ant{x, y, false, leaf{-1, -1, false}}
				tab[x][y] = "A"
				break
			}
		}
	}
	for i, _ := range tabLeafs {
		for {
			x := rand.Intn(10)
			y := rand.Intn(10)
			if tab[x][y] == "_" {
				tabLeafs[i] = leaf{x, y, false}
				tab[x][y] = "L"
				break
			}
		}
	}

	for i := 0; i < 100; i++ {
		time.Sleep(400 * time.Millisecond)
		// Czyszczenie konsoli, tylko dla Linux
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
		for _, strings := range tab {
			fmt.Println(strings)
		}
		for i, tabAnt := range tabAnts {
			for {
				addX := rand.Intn(3) - 1
				addY := rand.Intn(3) - 1
				if tabAnt.x+addX < 10 && tabAnt.y+addY < 10 && tabAnt.x+addX > 0 && tabAnt.y+addY > 0 && tab[tabAnt.x+addX][tabAnt.y+addY] == "_" {
					tab[tabAnt.x][tabAnt.y] = "_"
					tab[tabAnt.x+addX][tabAnt.y+addY] = "A"
					tabAnt.move(addX, addY)
					tabAnts[i] = tabAnt
					break
				} else if tabAnt.hasLeaf == false && tabAnt.x+addX < 10 && tabAnt.y+addY < 10 && tabAnt.x+addX > 0 && tabAnt.y+addY > 0 && tab[tabAnt.x+addX][tabAnt.y+addY] == "L" {
					tab[tabAnt.x+addX][tabAnt.y+addY] = "_"
					tabAnt.pickLeaf(findByCords(tabAnt.x+addX, tabAnt.y+addY, tabLeafs))
					tabAnts[i] = tabAnt
					break
				} else if tabAnt.hasLeaf == true && tabAnt.x+addX < 10 && tabAnt.y+addY < 10 && tabAnt.x+addX > 0 && tabAnt.y+addY > 0 && tab[tabAnt.x+addX][tabAnt.y+addY] == "L" {
					for {
						leftLeafX := rand.Intn(3) - 1
						leftLeafY := rand.Intn(3) - 1
						if tabAnt.x+leftLeafX < 10 && tabAnt.y+leftLeafY < 10 && tabAnt.x+leftLeafX > 0 && tabAnt.y+leftLeafY > 0 && tab[tabAnt.x+leftLeafX][tabAnt.y+leftLeafY] == "_" {
							tab[tabAnt.x+leftLeafX][tabAnt.y+leftLeafY] = "L"
							tabAnt.deleteLeaf(tabAnt.leaf, tabAnt.x+leftLeafX, tabAnt.y+leftLeafY)
							tabAnts[i] = tabAnt
							break
						}
					}
					break
				}
			}
		}
		fmt.Println()
	}
}
