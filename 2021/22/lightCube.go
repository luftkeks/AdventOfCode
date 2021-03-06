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
	countableAreas := &[]*Area{}

	for in, area := range areas {
		fmt.Println(in)
		noOverlap := true
		countableAreas, noOverlap = addIfPossible(countableAreas, &area)
		if noOverlap && area.on {
			countableAreasTemp := append((*countableAreas), area.copy())
			countableAreas = &countableAreasTemp
		}
	}

	counter2 := 0
	for _, area := range *countableAreas {
		counter2 += area.getNumberOfLit()
	}
	fmt.Printf("In Part Two are %v on.\n", counter2)
}

func addIfPossible(cAreas *[]*Area, area *Area) (*[]*Area, bool) {
	noOverlap := true
	countableAreas := *cAreas
	for ii, cArea := range countableAreas {
		overlap := area.getOverlap(cArea)
		if overlap.getNumberOfLit() != 0 {
			noOverlap = false
			if len(countableAreas) == 1 {
				countableAreas = []*Area{}
			} else if ii == 0 {
				countableAreas = countableAreas[1:]
			} else if ii == len(countableAreas)-1 {
				countableAreas = countableAreas[0 : len(countableAreas)-1]
			} else {
				countableAreas = append(countableAreas[0:ii], countableAreas[ii+1:]...)
			}
			if overlap.on {
				countableAreas = append(countableAreas, overlap.copy())
			}
			countableAreas = append(countableAreas, cArea.splitAreaAroundArea(overlap)...)
			for _, area1 := range area.splitAreaAroundArea(overlap) {
				countableAreasTemp, didNothing := addIfPossible(&countableAreas, area1)
				countableAreas = *countableAreasTemp
				if didNothing && area1.on {
					countableAreas = append(countableAreas, area1.copy())
				}
			}
			return &countableAreas, noOverlap

		}
	}
	return &countableAreas, noOverlap
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

func (a *Area) getOverlap(b *Area) (overlapp *Area) {
	return &Area{Xstart: Max(a.Xstart, b.Xstart), XFinish: Min(a.XFinish, b.XFinish), Ystart: Max(a.Ystart, b.Ystart), YFinish: Min(a.YFinish, b.YFinish), Zstart: Max(a.Zstart, b.Zstart), ZFinish: Min(a.ZFinish, b.ZFinish), on: a.on && b.on}
}

func (a *Area) copy() *Area {
	return &Area{Xstart: a.Xstart, XFinish: a.XFinish, Ystart: a.Ystart, YFinish: a.YFinish, Zstart: a.Zstart, ZFinish: a.ZFinish, on: a.on}
}

func (a *Area) splitAreaAroundArea(b *Area) []*Area {
	result := []*Area{}

	if a.Xstart < b.Xstart {
		xAreaStart := Area{Xstart: a.Xstart, XFinish: b.Xstart - 1, Ystart: a.Ystart, YFinish: a.YFinish, Zstart: a.Zstart, ZFinish: a.ZFinish, on: a.on}
		result = append(result, &xAreaStart)
	}

	if a.XFinish > b.XFinish {
		xAreaFinish := Area{Xstart: b.XFinish + 1, XFinish: a.XFinish, Ystart: a.Ystart, YFinish: a.YFinish, Zstart: a.Zstart, ZFinish: a.ZFinish, on: a.on}
		result = append(result, &xAreaFinish)
	}

	var xStart int
	if b.Xstart < a.Xstart {
		xStart = a.Xstart
	} else {
		xStart = b.Xstart
	}

	var xFinish int
	if b.XFinish > a.XFinish {
		xFinish = a.XFinish
	} else {
		xFinish = b.XFinish
	}

	if a.Ystart < b.Ystart {
		yAreaStart := Area{Xstart: xStart, XFinish: xFinish, Ystart: a.Ystart, YFinish: b.Ystart - 1, Zstart: a.Zstart, ZFinish: a.ZFinish, on: a.on}
		result = append(result, &yAreaStart)
	}

	if a.YFinish > b.YFinish {
		yAreaFinish := Area{Xstart: xStart, XFinish: xFinish, Ystart: b.YFinish + 1, YFinish: a.YFinish, Zstart: a.Zstart, ZFinish: a.ZFinish, on: a.on}
		result = append(result, &yAreaFinish)
	}

	var yStart int
	if b.Ystart < a.Ystart {
		yStart = a.Ystart
	} else {
		yStart = b.Ystart
	}

	var yFinish int
	if b.YFinish > a.YFinish {
		yFinish = a.YFinish
	} else {
		yFinish = b.YFinish
	}

	if a.Zstart < b.Zstart {
		zAreaStart := Area{Xstart: xStart, XFinish: xFinish, Ystart: yStart, YFinish: yFinish, Zstart: a.Zstart, ZFinish: b.Zstart - 1, on: a.on}
		result = append(result, &zAreaStart)
	}

	if a.ZFinish > b.ZFinish {
		zAreaFinish := Area{Xstart: xStart, XFinish: xFinish, Ystart: yStart, YFinish: yFinish, Zstart: b.ZFinish + 1, ZFinish: a.ZFinish, on: a.on}
		result = append(result, &zAreaFinish)
	}

	return result
}

func (a *Area) getNumberOfLit() int {
	result := 0
	xStuff := (a.XFinish - a.Xstart + 1)
	yStuff := (a.YFinish - a.Ystart + 1)
	zStuff := (a.ZFinish - a.Zstart + 1)
	if xStuff > 0 && yStuff > 0 && zStuff > 0 {
		if a.on {
			result = xStuff * yStuff * zStuff
		} else {
			result = -1
		}
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
