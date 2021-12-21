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
		panic("IRgendwas stimmt mit number2 nicht")
	}

	score1 := 0
	score2 := 0
	dice := DeterministicDice{}

	for {
		number1 = (number1-1+dice.get()+dice.get()+dice.get())%10 + 1
		score1 += number1

		if score1 >= 1000 {
			fmt.Printf("Player 1 wins. Winning Number: %v*%v= %v\n", score2, dice.timesRolled, score2*dice.timesRolled)
			break
		}

		number2 = (number2-1+dice.get()+dice.get()+dice.get())%10 + 1
		score2 += number2

		if score2 >= 1000 {
			fmt.Printf("Player 1 wins. Winning Number: %v*%v= %v\n", score1, dice.timesRolled, score1*dice.timesRolled)
			break
		}
	}

	score1 = 0
	score2 = 0
}
