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
}
