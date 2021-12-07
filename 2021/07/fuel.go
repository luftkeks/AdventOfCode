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
	numberOfSubs := make([]int, 2000)
	for scanner.Scan() {
		inString := scanner.Text()
		values := strings.Split(inString, ",")
		for _, number := range values {
			num, err := strconv.Atoi(number)
			if err != nil {
				panic("Some input not valid")
			}
			numberOfSubs[num]++
		}
	}

	minHeight := 2000
	minFuel := 999999
	for height, _ := range numberOfSubs {
		totalFuel := 0
		for height2, number := range numberOfSubs {
			fuel := Abs(height2 - height)
			totalFuel += fuel * number
		}
		if totalFuel < minFuel {
			minFuel = totalFuel
			minHeight = height
		}
	}

	fmt.Printf("The height with lowest fuel consumption for part one is: %v with %v fuel.\n", minHeight, minFuel)

	minHeight2 := 2000
	minFuel2 := 99999999999
	for jj := 0; jj < len(numberOfSubs); jj++ {
		totalFuel := 0
		for height2, number := range numberOfSubs {
			fuelDelta := Abs(height2 - jj)
				totalFuel += fuelDelta * (fuelDelta+1) / 2 * number
		}
		if totalFuel < minFuel2 {
			minFuel2 = totalFuel
			minHeight2 = jj
		}
	}
	fmt.Printf("The height with lowest fuel consumption for part two is: %v with %v fuel.\n", minHeight2, minFuel2)

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
