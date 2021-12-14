package main

import (
	"bufio"
	"fmt"
	"os"
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
	transitions := map[string][]string{}
	scanner.Scan()
	inputString := scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		inString := scanner.Text()
		split := strings.Split(inString, " -> ")
		transitions[split[0]] = []string{fmt.Sprintf("%v%v", string(split[0][0]), split[1]), fmt.Sprintf("%v%v", split[1], string(split[0][1]))}
	}
	fmt.Printf("Input string was %v\n", inputString)
	working := map[string]int{}
	for jj := 0; jj < len(inputString)-1; jj++ {
		working[inputString[jj:jj+2]]++
	}
	for ii := 0; ii < 10; ii++ {
		maap := map[string]int{}
		for key, value := range working {
			trans, ok := transitions[key]
			if ok {
				maap[trans[0]] += value
				maap[trans[1]] += value
			}
		}
		working = copyMap(maap)
	}

	counting := map[rune]int{rune(inputString[len(inputString)-1]): 1}
	for key, value := range working {
		counting[rune(key[1])] += value
	}
	fmt.Println(counting)
	max := 0
	min := counting['N']
	for _, value := range counting {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	fmt.Printf("The Result of one is: %v and %v which results to: %v\n", max, min, max-min)

	for ii := 0; ii < 30; ii++ {
		maap := map[string]int{}
		for key, value := range working {
			trans, ok := transitions[key]
			if ok {
				maap[trans[0]] += value
				maap[trans[1]] += value
			}
		}
		working = copyMap(maap)
	}

	counting = map[rune]int{}
	for key, value := range working {
		counting[rune(key[1])] += value
	}
	fmt.Println(counting)
	max = 0
	min = counting['N']
	for _, value := range counting {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	fmt.Printf("The Result of two is: %v and %v which results to: %v\n", max, min, max-min)
}

func copyMap(maap map[string]int) map[string]int {
	result := map[string]int{}
	for key, value := range maap {
		result[key] = value
	}
	return result
}
