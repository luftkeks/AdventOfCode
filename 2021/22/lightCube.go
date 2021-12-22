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

type Position struct {
	x, y, z int
}

type Area struct {
	Xstart, XFinish int
	Ystart, YFinish int
	Zstart, ZFinish int
	on              bool
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

	areas := readInLines(lines)

	// Part 1
	dots1 := map[Position]bool{}
	for _, area := range areas {
		if area.inBounds(50) {
			for xx := area.Xstart; xx <= area.XFinish; xx++ {
				for yy := area.Ystart; yy <= area.YFinish; yy++ {
					for zz := area.Zstart; zz <= area.ZFinish; zz++ {
						dots1[Position{x: xx, y: yy, z: zz}] = area.on
					}
				}
			}
		}
	}

	counter1 := 0
	for _, value := range dots1 {
		if value {
			counter1++
		}
	}

	fmt.Printf("In Part One are %v on.\n", counter1)

	// Part 2
	dots2 := map[Position]bool{}
	for _, area := range areas {
		// This has to be recursive areas with sub areas which then difine on and off parts and kann be just added and subtracted
	}

	counter2 := 0
	for _, value := range dots2 {
		if value {
			counter2++
		}
	}

	fmt.Printf("In Part Two are %v on.\n", counter2)
}

func readInLines(lines []string) []Area {
	areas := []Area{}
	for _, line := range lines {
		inStuff := strings.Split(line, " ")
		xyz := strings.Split(inStuff[1], ",")

		on := inStuff[0] == "on"
		xSplit := strings.Split(xyz[0], "..")
		xStart, _ := strconv.Atoi(xSplit[0][2:])
		xfinish, _ := strconv.Atoi(xSplit[1])
		ySplit := strings.Split(xyz[1], "..")
		yStart, _ := strconv.Atoi(ySplit[0][2:])
		yFinish, _ := strconv.Atoi(ySplit[1])
		zSplit := strings.Split(xyz[2], "..")
		zStart, _ := strconv.Atoi(zSplit[0][2:])
		zFinish, _ := strconv.Atoi(zSplit[1])
		areas = append(areas, Area{Xstart: xStart, XFinish: xfinish, Ystart: yStart, YFinish: yFinish, Zstart: zStart, ZFinish: zFinish, on: on})
	}
	return areas
}

func (a *Area) inBounds(border int) bool {
	return Abs(a.Xstart) <= border && Abs(a.XFinish) <= border && Abs(a.Ystart) <= border && Abs(a.YFinish) <= border && Abs(a.Zstart) <= border && Abs(a.ZFinish) <= border
}

// TODO implement this
func (a *Area) getOverlap(b *Area) (in, out Area) {
	return 0, 0
}

func (a *Area) getNumberOfLit() int {

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
