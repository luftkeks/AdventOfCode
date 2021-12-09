package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type point struct {
	xx, yy int
	height rune
}

type basin struct {
	points []point
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

			dot := point{xx: xx, yy: yy, height: rune(lines[yy][xx])}

			if dot.height < '9' {
				hasBasin := false
				bas := -1
				for ii := 0; ii < len(basins); ii++ {
					if hasBasin {
						if isPointAdjectedToBasin(dot, basins[ii]) {
							basins[ii].points = append(basins[ii].points, basins[bas].points...)
							basins[bas].points = []point{}
						}
						continue
					}

					isInBasin := isPointAdjectedToBasin(dot, basins[ii])
					if isInBasin {
						basins[ii].points = append(basins[ii].points, dot)
						bas = ii
					}
					hasBasin = isInBasin
				}
				if !hasBasin {
					dots := []point{dot}
					basins = append(basins, basin{points: dots})
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
		riskLevel += uint(pointt.height) - uint('0') + 1
	}

	fmt.Printf("The risk level of the lowest points in the map is: %v\n", riskLevel)

	biggest := make([]int, 3)
	indexSmallest := 0
	for _, bas := range basins {
		lang := len(bas.points)
		if lang > biggest[indexSmallest] {
			biggest[indexSmallest] = lang
			indexSmallest = getSmallestIndest(biggest)
		}
	}

	fmt.Printf("The 3 biggest basins multiplied by their length are: %v\n", biggest[0]*biggest[1]*biggest[2])
}

func (b *basin) checkPoint(dot point) bool {
	for _, thing := range b.points {
		if isPointAdjected(thing, dot) {
			return true
		}
	}
	return false
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

func isPointInSlice(list []point, dot point) bool {
	for _, element := range list {
		if element == dot {
			return true
		}
	}
	return false
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

func isPointAdjectedToBasin(dot point, bas basin) bool {
	for _, basinDot := range bas.points {
		if isPointAdjected(dot, basinDot) {
			return true
		}
	}
	return false
}

func getSmallestIndest(in []int) int {
	indexSmallest := 0
	smallest := 999999999
	for index, value := range in {
		if value < smallest {
			indexSmallest = index
		}
	}
	return indexSmallest
}
