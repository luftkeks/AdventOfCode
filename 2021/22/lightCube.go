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
	partOne(areas)

	// Part 2
	countableAreas := []Area{}

	for _, area := range areas {

	}

	counter2 := 0
	for _, area := range countableAreas {
		counter2 += area.getNumberOfLit()
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

func (a *Area) getOverlap(b *Area) (overlapp Area) {
	return Area{Xstart: Max(a.Xstart, b.Xstart), XFinish: Min(a.XFinish, b.XFinish), Ystart: Max(a.Ystart, b.Ystart), YFinish: Min(a.YFinish, b.YFinish), Zstart: Max(a.Zstart, b.Zstart), ZFinish: Min(a.ZFinish, b.ZFinish), on: a.on && b.on}
}

func (a *Area) overlapps(b *Area) bool {
	if a.Xstart <= b.XFinish && a.XFinish >= b.XFinish {
		return true
	}
	if a.Xstart <= b.Xstart && a.XFinish >= b.Xstart {
		return true
	}
	if a.Ystart <= b.Ystart && a.YFinish >= b.Ystart {
		return true
	}
	if a.Ystart <= b.YFinish && a.YFinish >= b.YFinish {
		return true
	}
	if a.Zstart <= b.Zstart && a.ZFinish >= b.Zstart {
		return true
	}
	if a.Zstart <= b.ZFinish && a.ZFinish >= b.ZFinish {
		return true
	}
	return false
}

func createMatchingSubCubes(area1, area2 Area) []Area {
	// this has to be implemented
	// Check for borderes and create an overlapping area
	// if both are lit keep overlapp once if only one is lit keep area without overlapp
	// make sub cubes around the overlapping area
	// if there is a lit area and an off area keep none of them - rest keep only lit
	result := []Area{}
	return result
}

func (a *Area) getNumberOfLit() int {
	result := 0
	if a.on {
		result = Abs(a.XFinish-a.Xstart+1) * Abs(a.YFinish-a.Ystart+1) * Abs(a.ZFinish-a.Zstart+1)
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func partOne(areas []Area) {
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
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
