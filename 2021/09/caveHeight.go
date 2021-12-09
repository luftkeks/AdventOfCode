package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type point struct {
	xx, yy int
	height int
}

type basin struct {
	points    []point
	lastPoint point
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

	lowestPoints := []point{}
	basins := []basin{}
	for yy := 0; yy < len(lines); yy++ {
		for xx := 0; xx < len(lines[yy]); xx++ {

			dot := point{xx: xx, yy: yy, height: int(lines[yy][xx]) - 0x30}

			if dot.height < 9 {
				hasBasin := false
				lastMatchingBasin := -1
				for ii := 0; ii < len(basins); ii++ {
					if hasBasin {
						if basins[ii].isPointAdjectedToBasin(dot) {
							basins[ii].points = append(basins[ii].points, basins[lastMatchingBasin].points...)
							basins[lastMatchingBasin].points = []point{}
							lastMatchingBasin = ii
						}
						continue
					}

					isInBasin := basins[ii].isPointAdjectedToBasin(dot)
					if isInBasin {
						basins[ii].addPoint(dot)
						lastMatchingBasin = ii
					}
					hasBasin = isInBasin
				}
				if !hasBasin {
					newBasin := basin{}
					newBasin.addPoint(dot)
					basins = append(basins, newBasin)
				}
			}

			if numberNotOverEdge(lines, yy, xx-1) && lines[yy][xx-1] <= lines[yy][xx] {
				continue
			}
			if numberNotOverEdge(lines, yy, xx+1) && lines[yy][xx+1] <= lines[yy][xx] {
				continue
			}
			if numberNotOverEdge(lines, yy+1, xx) && lines[yy+1][xx] <= lines[yy][xx] {
				continue
			}
			if numberNotOverEdge(lines, yy-1, xx) && lines[yy-1][xx] <= lines[yy][xx] {
				continue
			}

			lowestPoints = append(lowestPoints, dot)
		}
	}

	riskLevel := uint(0)
	for _, pointt := range lowestPoints {
		riskLevel += uint(pointt.height) + 1
	}

	fmt.Printf("The risk level of the lowest points in the map is: %v\n", riskLevel)

	biggest := make([]int, 3)
	indexSmallest := 0
	for _, bas := range basins {
		lang := len(bas.points)
		if lang >= biggest[indexSmallest] {
			biggest[indexSmallest] = lang
			indexSmallest = getSmallestIndex(biggest)
		}
	}

	fmt.Printf("The 3 biggest basins multiplied by their length are: %v\n", biggest[0]*biggest[1]*biggest[2])
}

func numberNotOverEdge(maap []string, yy, xx int) bool {
	yBorder := len(maap)
	xBorder := len(maap[0])
	if !(xx >= 0 && xx < xBorder) {
		return false
	} else if !(yy >= 0 && yy < yBorder) {
		return false
	}
	return true
}

func isPointAdjected(point1, point2 point) bool {
	if Abs(point1.xx-point2.xx)+Abs(point1.yy-point2.yy) == 1 {
		return true
	} else {
		return false
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (b *basin) isPointAdjectedToBasin(dot point) bool {
	if isPointAdjected(dot, b.lastPoint) {
		return true
	}
	for ii := 0; ii < len(b.points); ii++ {
		if isPointAdjected(dot, b.points[ii]) {
			return true
		}
	}
	return false
}

func (b *basin) addPoint(dot point) {
	b.lastPoint = dot
	b.points = append(b.points, dot)
}

func getSmallestIndex(in []int) int {
	indexSmallest := 0
	smallest := 9999
	for index, value := range in {
		if value < smallest {
			indexSmallest = index
			smallest = value
		}
	}
	return indexSmallest
}
