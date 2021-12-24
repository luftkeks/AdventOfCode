package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}

func main() {
	defer elapsed()()
	dat, err := os.Open("input.txt")
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

	multimap := map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}
	player1, player2 := 0, 0
	for winCond := 0; winCond < 22; winCond++ {
		player1, player2 = startRound(winCond, number1, number2, 0, 0, 1, &multimap)
		fmt.Printf("times won winCond: %v player1: %v player2: %v\n", winCond, player1, player2)
	}

	fmt.Printf(" vgl mit lösung 1: %v       2: %v\n", player1-444356092776315, player2-341960390180808)
}

func startRound(winCond, number1old, number2old, score1old, score2old int, multiplicator int, multimap *map[int]int) (int, int) {
	winCondition := winCond
	player1won := 0
	player2won := 0
	for dice1 := 3; dice1 <= 9; dice1++ {
		number1new, score1new, multi1New := doStep(dice1, number1old, score1old, multiplicator, multimap)
		if score1new >= winCondition {
			player1won += multi1New
			continue
		}
		for dice2 := 3; dice2 <= 9; dice2++ {
			number2new, score2new, multi2New := doStep(dice2, number2old, score2old, multi1New, multimap)
			if score2new >= winCondition {
				player2won += multi2New
			} else {
				player1wonTemp, player2wonTemp := startRound(winCond, number1new, number2new, score1new, score2new, multi2New, multimap)
				player1won += player1wonTemp
				player2won += player2wonTemp
			}
		}
	}
	return player1won, player2won
}

func doStep(dice, number, score, multi int, multimap *map[int]int) (numberNew, scoreNew, multiNew int) {
	numberNew = (number-1+dice)%10 + 1
	scoreNew = score + numberNew
	multiNew = multi * (*multimap)[dice]
	return
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

type DeterministicDice struct {
	number, timesRolled int
}

func (d *DeterministicDice) get() int {
	d.number = (d.number)%100 + 1
	d.timesRolled++
	return d.number
}
