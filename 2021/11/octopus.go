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
	lines := []string{}
	for scanner.Scan() {
		inString := scanner.Text()
		lines = append(lines, inString)
	}

	maap := make([][]int, len(lines))
	for in, line := range lines {
		maap[in] = make([]int, len(line))
		for in2, run := range line {
			maap[in][in2] = int(run - '0')
		}
	}

	couter := 0
	flash := 0
	all := false
	for ii := 0; ii < 1000; ii++ {
		maap, flash, all = step(maap)
		couter += flash
		if all {
			fmt.Printf("The Step in which all elight is: %v\n", ii+1)
			break
		}
		if ii == 99 {
			fmt.Printf("The number of flashes after Part 1 is: %v\n", couter)
		}
	}

}

func step(maap [][]int) ([][]int, int, bool) {
	size := len(maap) * len(maap[0])
	for yy := 0; yy < len(maap); yy++ {
		for xx := 0; xx < len(maap[0]); xx++ {
			maap[yy][xx]++
		}
	}
	coutner := 0
	for {
		flash := 0
		workingMaap := append([][]int{}, maap...)
		for yy := 0; yy < len(maap); yy++ {
			for xx := 0; xx < len(maap[0]); xx++ {
				if maap[yy][xx] > 9 {
					flash++
					for ii := -1; ii <= 1; ii++ {
						for jj := -1; jj <= 1; jj++ {
							if numberNotOverEdge(maap, yy+ii, xx+jj) {
								workingMaap[yy+ii][xx+jj]++
							}
						}
					}
					workingMaap[yy][xx] = -9999
				}
			}
		}
		if flash == 0 {
			break
		} else {
			coutner += flash
			maap = workingMaap
		}
	}
	for yy := 0; yy < len(maap); yy++ {
		for xx := 0; xx < len(maap[0]); xx++ {
			if maap[yy][xx] < 0 {
				maap[yy][xx] = 0
			}
		}
	}
	return maap, coutner, coutner == size
}

func numberNotOverEdge(maap [][]int, yy, xx int) bool {
	yBorder := len(maap)
	xBorder := len(maap[0])
	if !(xx >= 0 && xx < xBorder) {
		return false
	} else if !(yy >= 0 && yy < yBorder) {
		return false
	}
	return true
}

func print(maap [][]int) {
	for _, row := range maap {
		for _, number := range row {
			fmt.Printf("%v", number)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}
