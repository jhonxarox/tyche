package main

import (
	"fmt"
	"math/rand"
)

const (
	round = 4
)

type Player struct {
	name  string
	value int
}

//random number function
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

//find max or the winner
func winner(players []Player) Player {
	var winner Player
	var max int = 0
	for i := 0; i < len(players); i++ {
		if players[i].value > max {
			winner = players[i]
		}
	}
	return winner
}

func main() {
	var playersNumber int

	fmt.Println("Insert Player Number :")
	fmt.Scan(&playersNumber)

	var players = make([]Player, playersNumber)
	var name string

	//add player name
	for i := 0; i < len(players); i++ {
		fmt.Println("Player Name :")
		fmt.Scan(&name)
		players[i] = Player{name, 0}
	}

	for i := 0; i < round; i++ {
		fmt.Println("Round", i+1)
		for j := 0; j < len(players); j++ {
			randomNumber := random(1, 6)
			fmt.Println("Player Name : ", players[j].name)
			fmt.Println("Player Current Value : ", players[j].value)
			fmt.Println("Dice Number : ", randomNumber)

			if randomNumber%2 == 1 {
				players[j].value += 10
			} else {
				players[j].value -= 5
			}
		}
	}
	fmt.Println("Winner : ", winner(players))
}
