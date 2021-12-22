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

func playRound(dice1, dice2, number1, number2, score1, score2 int, multiplicator int) (player1won, player2won int) {
	winCondition := 21
	multi1 := multimap[dice1]
	multi2 := multimap[dice2]

	number1 = (number1-1+dice1)%10 + 1
	number2 = (number2-1+dice2)%10 + 1
	score1 += number1
	score2 += number2

	if score1 >= winCondition {
		return multiplicator * multi1, 0
	}
	if score2 >= winCondition {
		return 0, multiplicator * multi1 * multi2
	}
	return startRound(number1, number2, score1, score2, multiplicator*multi1*multi2)
}

func startRound(number1, number2, score1, score2 int, multiplicator int) (int, int) {
	player1won := 0
	player2won := 0
	for dice1 := 3; dice1 <= 9; dice1++ {
		for dice2 := 3; dice2 <= 9; dice2++ {
			player1wonTemp, player2wonTemp := playRound(dice1, dice2, number1, number2, score1, score2, multiplicator)
			player1won += player1wonTemp
			player2won += player2wonTemp
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
