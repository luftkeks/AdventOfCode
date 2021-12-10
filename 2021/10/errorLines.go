package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

//var legalChars = []rune{'(', '[', '{', '<'}
var closingChars = map[rune]bool{')': true, ']': true, '}': true, '>': true}
var legalMap = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
var illegalCharPointMap = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var closingCharPointMap = map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}

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

	one, two := solutionPartOneAndTwo(lines)
	fmt.Printf("The Solution of Part One is: %v\n", one)
	fmt.Printf("The Solution of Part One is: %v\n", two)
}

func solutionPartOneAndTwo(lines []string) (int, int) {
	pointsPartOne := 0
	pointsPartTwo := []int{}
	for _, line := range lines {
		char, point2 := solveLine(line)
		if !(char == '-' || char == '_') {
			pointsPartOne += illegalCharPointMap[char]
		}
		if point2 != 0 {
			pointsPartTwo = append(pointsPartTwo, point2)
		}
	}
	sort.Ints(pointsPartTwo)
	return pointsPartOne, pointsPartTwo[len(pointsPartTwo)/2]
}

func solveLine(line string) (rune, int) {
	for ii := 0; ii < len(line)-1; ii++ {
		// I don't know why i need this but i need this
		if isMatchingChars(rune(line[0]), rune(line[1])) {
			line = line[2:]
			ii = 0
		}
		if isMatchingChars(rune(line[ii]), rune(line[ii+1])) {
			line = line[:ii] + line[ii+2:]
			ii = 0
		}
	}
	if len(line) == 0 {
		return '-', 0
	} else {
		fcc := firstClosingChar(line)
		if fcc != '_' {
			return fcc, 0
		} else {
			number := solvePartTwo(line)
			return fcc, number
		}
	}
}

func isMatchingChars(one, two rune) bool {
	return two == legalMap[one]
}

func firstClosingChar(line string) rune {
	for _, char := range line {
		if closingChars[char] {
			return char
		}
	}
	return '_'
}

func solvePartTwo(incompleteLine string) int {
	points := 0
	for ii := len(incompleteLine) - 1; ii >= 0; ii-- {
		char := rune(incompleteLine[ii])
		points = points*5 + closingCharPointMap[char]
	}
	return points
}
