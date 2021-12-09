package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

type point struct {
	xx, yy, height int
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
	dat, err := os.Open("test.txt")
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
	for index := range heightMap {
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
	basins := []basin{}
	for yy := 0; yy < len(heightMap); yy++ {
		for xx := 0; xx < len(heightMap[yy]); xx++ {

			dot := point{xx: xx, yy: yy, height: heightMap[yy][xx]}

			if dot.height < 9 {
				hasBasin := false
				bas := -1
				for ii := 0; ii < len(basins); ii++ {
					if hasBasin {
						if isPointAdjectedToBasin(dot, basins[ii]) {
							basins[ii].points = append(basins[ii].points, basins[bas].points...)
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

			lowestPoints = append(lowestPoints, dot)
		}
	}

	riskLevel := 0
	for _, pointt := range lowestPoints {
		riskLevel += pointt.height + 1
	}

	fmt.Printf("The risk level of the lowest points in the map is: %v\n", riskLevel)

	// cleanup Basins
	resultBasins := []basin{}
	for _, bas := range basins {
		isInResult := false
		for _, resBas := range resultBasins {
			if resBas.checkPoint(bas.points[0]) {
				isInResult = true
			}
		}
		if !isInResult {
			resultBasins = append(resultBasins, bas)
		}
	}

	sort.SliceStable(resultBasins, func(i, j int) bool { return len(resultBasins[i].points) > len(resultBasins[j].points) })

	fmt.Println(resultBasins)
}

func (b *basin) checkPoint(dot point) bool {
	for _, thing := range b.points {
		if isPointAdjected(thing, dot) {
			return true
		}
	}
	return false
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
