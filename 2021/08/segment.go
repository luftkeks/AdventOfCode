package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	counter := 0
	for scanner.Scan() {
		inString := scanner.Text()
		lines = append(lines, inString)
		inOut := strings.Split(inString, "|")
		outs := strings.Split(inOut[1], " ")
		for _, symbol := range outs {
			if len(symbol) == 2 || len(symbol) == 3 || len(symbol) == 4 || len(symbol) == 7 {
				counter++
			}
		}
	}

	fmt.Printf("The number of 1478 is %v.\n", counter)

	counter2 := 0
	for _, line := range lines {
		inOut := strings.Split(line, "|")
		digitsIn := strings.Split(inOut[0], " ")
		digitsOut := strings.Split(inOut[1], " ")
		solv := createSolver()
		sort.SliceStable(digitsIn, func(i, j int) bool { return len(digitsIn[i]) < len(digitsIn[j]) })
		for _, digit := range digitsIn {
			switch len(digit) {
			case 2:
				solv.addToMap(digit, 1)
			case 3:
				solv.addToMap(digit, 7)
			case 4:
				solv.addToMap(digit, 4)
			case 5:
				if strings.ContainsRune(digit, solv.haken[0]) && strings.ContainsRune(digit, solv.haken[1]) {
					solv.addToMap(digit, 5)
				} else if containsRunesOfString(digit, solv.mappe[1]) {
					solv.addToMap(digit, 3)
				} else {
					solv.addToMap(digit, 2)
				}
			case 7:
				solv.addToMap(digit, 8)
			case 6:
				if !(containsRunesOfString(digit, solv.mappe[1])) {
					solv.addToMap(digit, 6)
				} else if containsRunesOfString(digit, solv.mappe[4]) {
					solv.addToMap(digit, 9)
				} else {
					solv.addToMap(digit, 0)
				}
			}
		}

		zahl := make([]int, 5)
		for stelle, blub := range digitsOut {
			for key, value := range solv.mappe {
				if containsRunesOfString(value, blub) && containsRunesOfString(blub, value) {
					zahl[stelle-1] = key
					continue
				}
			}
		}

		zahlString := fmt.Sprintf("%v%v%v%v", zahl[0], zahl[1], zahl[2], zahl[3])
		number, err := strconv.Atoi(zahlString)
		if err != nil {
			panic("Kaputte zahl")
		}
		counter2 += number
	}

	fmt.Printf("The sum of all is %v.\n", counter2)
}

type solver struct {
	mappe map[int]string
	haken []rune
}

func createSolver() solver {
	solv := solver{mappe: map[int]string{}, haken: []rune{}}
	return solv
}

func (s *solver) addToMap(str string, number int) {
	s.mappe[number] = str
	if number == 4 {
		for _, runE := range str {
			if !(strings.ContainsRune(s.mappe[1], runE)) {
				s.haken = append(s.haken, runE)
			}
		}
	}
}

func containsRunesOfString(in, runes string) bool {
	result := true
	for _, runE := range runes {
		if result {
			result = strings.ContainsRune(in, runE)
		}
	}
	return result
}
