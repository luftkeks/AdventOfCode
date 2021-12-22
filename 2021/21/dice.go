package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var multimap map[int]int

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}

type DeterministicDice struct {
	number, timesRolled int
}

func (d *DeterministicDice) get() int {
	d.number = (d.number)%100 + 1
	d.timesRolled++
	return d.number
}

func main() {
	defer elapsed()()
	dat, err := os.Open("test.txt")
	if err != nil {
		panic("Hilfe File tut nicht")
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	lines := []string{}
	for scanner.Scan() {
		inString := scanner.Text()
		lines = append(lines, inString)
	}

	line1 := strings.Split(lines[0], " ")
	line2 := strings.Split(lines[1], " ")

	number1, err := strconv.Atoi(line1[len(line1)-1])
	if err != nil {
		panic("IRgendwas stimmt mit number1 nicht")
	}
	number2, err := strconv.Atoi(line2[len(line2)-1])
	if err != nil {
		panic("Irgendwas stimmt mit number2 nicht")
	}

	partOne(number1, number2)

	multimap = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}
	player1, player2 := startRound(number1, number2, 0, 0, 1)

	fmt.Printf("times won player1: %v player2: %v\n", player1, player2)
	fmt.Printf(" vgl mit lösung 1: %v       2: %v\n", player1-444356092776315, player2-341960390180808)
}

func startRound(number1old, number2old, score1old, score2old int, multiplicator int) (int, int) {
	winCondition := 21
	player1won := 0
	player2won := 0
	for dice1 := 3; dice1 <= 9; dice1++ {
		for dice2 := 3; dice2 <= 9; dice2++ {
			number1new := (number1old-1+dice1)%10 + 1
			number2new := (number2old-1+dice2)%10 + 1
			score1new := score1old + number1new
			score2new := score2old + number2new
			if score1new >= winCondition {
				player1won += multiplicator * multimap[dice1]
			} else if score2new >= winCondition {
				player2won += multiplicator * multimap[dice1] * multimap[dice2]
			} else {
				player1wonTemp, player2wonTemp := startRound(number1new, number2new, score1new, score2new, multiplicator*multimap[dice1]*multimap[dice2])
				player1won += player1wonTemp
				player2won += player2wonTemp
			}
		}
	}
	return player1won, player2won
}

func partOne(number1, number2 int) {
	score1 := 0
	score2 := 0
	dice := DeterministicDice{}

	for {
		number1 = (number1-1+dice.get()+dice.get()+dice.get())%10 + 1
		score1 += number1

		if score1 >= 1000 {
			fmt.Printf("Player 1 wins. Winning Number: %v*%v = %v\n", score2, dice.timesRolled, score2*dice.timesRolled)
			return
		}

		number2 = (number2-1+dice.get()+dice.get()+dice.get())%10 + 1
		score2 += number2

		if score2 >= 1000 {
			fmt.Printf("Player 1 wins. Winning Number: %v*%v = %v\n", score1, dice.timesRolled, score1*dice.timesRolled)
			return
		}
	}
}
