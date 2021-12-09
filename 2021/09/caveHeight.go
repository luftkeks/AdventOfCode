package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type point struct {
	xx, yy, height int
}

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

	heightMap := make([][]int, len(lines))
	for index, _ := range heightMap {
		heightMap[index] = make([]int, len(lines[index]))
		for indexLine, char := range lines[index] {
			number, err := strconv.Atoi(string(char))
			if err != nil {
				panic("Something doesnt work with the number conversion of " + string(char))
			}
			heightMap[index][indexLine] = number
		}
	}

	lowestPoints := []point{}
	for yy := 0; yy < len(heightMap); yy++ {
		for xx := 0; xx < len(heightMap[yy]); xx++ {
			if numberNotOverEdge(heightMap, yy, xx-1) && heightMap[yy][xx-1] <= heightMap[yy][xx] {
				continue
			}
			if numberNotOverEdge(heightMap, yy, xx+1) && heightMap[yy][xx+1] <= heightMap[yy][xx] {
				continue
			}
			if numberNotOverEdge(heightMap, yy+1, xx) && heightMap[yy+1][xx] <= heightMap[yy][xx] {
				continue
			}
			if numberNotOverEdge(heightMap, yy-1, xx) && heightMap[yy-1][xx] <= heightMap[yy][xx] {
				continue
			}

			lowestPoints = append(lowestPoints, point{xx: xx, yy: yy, height: heightMap[yy][xx]})
		}
	}

	riskLevel := 0
	for _, pointt := range lowestPoints {
		riskLevel += pointt.height + 1
	}

	fmt.Printf("The risk level of the lowest points in the map is: %v\n", riskLevel)

	for _, dot := range lowestPoints {
		basin := []point{}

	}
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

func isPointInStruct(list []point, dot point) bool {
	for _, element := range list {
		if element == dot {
			return true
		}
	}
	return false
}
