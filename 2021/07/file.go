package main

import (
	"bufio"
	"fmt"
	"os"
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
	input := []string{}
	for scanner.Scan() {
		inString := scanner.Text()
		input = append(input, inString)
	}

	fmt.Println(input)
}
