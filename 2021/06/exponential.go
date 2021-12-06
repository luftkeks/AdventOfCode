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
	dat, _ := os.Open("input.txt")
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	daysTilBirth := make([]int, 10)
	for scanner.Scan() {
		inString := scanner.Text()
		values := strings.Split(inString, ",")
		for _, number := range values {
			num, err := strconv.Atoi(number)
			if err != nil {
				panic("Some input not valid")
			}
			daysTilBirth[num]++
		}
	}

	for ii := 0; ii < 80; ii++ {
		newFishes := daysTilBirth[0]
		for jj := 0; jj < 9; jj++ {
			daysTilBirth[jj] = daysTilBirth[jj+1]
		}
		daysTilBirth[6] += newFishes
		daysTilBirth[8] += newFishes
	}

	numberOfFishes := 0
	for _, fish := range daysTilBirth {
		numberOfFishes += fish
	}

	fmt.Printf("There are %v fishes after 80 Days\n", numberOfFishes)

	for ii := 80; ii < 256; ii++ {
		newFishes := daysTilBirth[0]
		for jj := 0; jj < 9; jj++ {
			daysTilBirth[jj] = daysTilBirth[jj+1]
		}
		daysTilBirth[6] += newFishes
		daysTilBirth[8] += newFishes
	}

	numberOfFishes = 0
	for _, fish := range daysTilBirth {
		numberOfFishes += fish
	}
	fmt.Printf("There are %v fishes after 256 Days\n", numberOfFishes)
}
