package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GridBool struct {
	lamps [][]bool
}

type GridBright struct {
	lamps [][]int
}

type Grid interface {
	CountLit() int
	Switch(on bool, xStart, yStart, xEnd, yEnd int)
	Toggle(xStart, yStart, xEnd, yEnd int)
}

func main() {
	commandLineArgs := os.Args
	fileToRead := commandLineArgs[1]
	dat, _ := os.Open(fileToRead)
	scanner := bufio.NewScanner(dat)
	scannedStrings := []string{}
	for scanner.Scan() {
		scannedStrings = append(scannedStrings, scanner.Text())
	}

	grid := make([][]bool, 1000)
	for ii := range grid {
		grid[ii] = make([]bool, 1000)
		for jj := range grid[ii] {
			grid[ii][jj] = false
		}
	}
	lampions := GridBool{lamps: grid}

	grid2 := make([][]int, 1000)
	for ii := range grid2 {
		grid2[ii] = make([]int, 1000)
		for jj := range grid2[ii] {
			grid2[ii][jj] = 0
		}
	}
	lampions2 := GridBright{lamps: grid2}

	for _, line := range scannedStrings {
		ParseLine(lampions, line)
		ParseLine(lampions2, line)
	}

	fmt.Println("Number of lit lamps: ", lampions.CountLit())
	fmt.Println("Total number of brightnes: ", lampions2.CountLit())

}

func ParseLine(grid Grid, line string) { // turn on 0,0 through 999,999
	words := strings.Split(line, " ")
	if words[0] == "turn" {
		start := strings.Split(words[2], ",")
		end := strings.Split(words[4], ",")
		xStart, _ := strconv.Atoi(start[0])
		yStart, _ := strconv.Atoi(start[1])
		xEnd, _ := strconv.Atoi(end[0])
		yEnd, _ := strconv.Atoi(end[1])
		grid.Switch(words[1] == "on", xStart, yStart, xEnd, yEnd)
	} else if words[0] == "toggle" {
		start := strings.Split(words[1], ",")
		end := strings.Split(words[3], ",")
		xStart, _ := strconv.Atoi(start[0])
		yStart, _ := strconv.Atoi(start[1])
		xEnd, _ := strconv.Atoi(end[0])
		yEnd, _ := strconv.Atoi(end[1])
		grid.Toggle(xStart, yStart, xEnd, yEnd)
	}
}

func (g GridBool) CountLit() int {
	lit := 0
	for _, row := range g.lamps {
		for _, lamp := range row {
			if lamp {
				lit++
			}
		}
	}
	return lit
}

func (g GridBool) Switch(on bool, xStart, yStart, xEnd, yEnd int) {
	for ii := xStart; ii <= xEnd; ii++ {
		for jj := yStart; jj <= yEnd; jj++ {
			g.lamps[ii][jj] = on
		}
	}
}

func (g GridBool) Toggle(xStart, yStart, xEnd, yEnd int) {
	for ii := xStart; ii <= xEnd; ii++ {
		for jj := yStart; jj <= yEnd; jj++ {
			g.lamps[ii][jj] = !g.lamps[ii][jj]
		}
	}
}

func (g GridBright) CountLit() int {
	lit := 0
	for _, row := range g.lamps {
		for _, lamp := range row {
			lit += lamp
		}
	}
	return lit
}

func (g GridBright) Switch(on bool, xStart, yStart, xEnd, yEnd int) {
	for ii := xStart; ii <= xEnd; ii++ {
		for jj := yStart; jj <= yEnd; jj++ {
			if on {
				g.lamps[ii][jj] += 1
			} else if g.lamps[ii][jj] > 0 {
				g.lamps[ii][jj] -= 1
			}
		}
	}
}

func (g GridBright) Toggle(xStart, yStart, xEnd, yEnd int) {
	for ii := xStart; ii <= xEnd; ii++ {
		for jj := yStart; jj <= yEnd; jj++ {
			g.lamps[ii][jj] += 2
		}
	}
}
