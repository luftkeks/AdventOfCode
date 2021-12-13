package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer elapsed()()

	dots := len(parseWithFoldings("input.txt", 1))
	fmt.Printf("The Answer to Part one is: %v\n", dots)

	points := parseWithFoldings("input.txt", 100)
	xMax, yMax := Max(points)
	lines := make([][]rune, yMax+1)
	for ii := 0; ii < len(lines); ii++ {
		line := []rune{}
		for jj := 0; jj < xMax+1; jj++ {
			line = append(line, ' ')
		}
		lines[ii] = line
	}
	for _, dot := range points {
		lines[dot.y][dot.x] = '#'
	}
	str := []string{}
	for _, line := range lines {
		str = append(str, string(line)+"\n")
	}
	fmt.Println(str)
}

func parseWithFoldings(input string, folding int) []Point {
	dat, err := os.Open(input)
	if err != nil {
		panic("Hilfe File tut nicht")
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	points := []Point{}
	foldings := []Folding{}
	for scanner.Scan() {
		inString := scanner.Text()
		split := strings.Split(inString, ",")
		if len(split) == 2 {
			xx, _ := strconv.Atoi(split[0])
			yy, _ := strconv.Atoi(split[1])
			points = append(points, Point{x: xx, y: yy})
		} else if len(inString) < 5 {
			continue
		} else {
			split2 := strings.Split(inString, " ")
			split3 := strings.Split(split2[2], "=")
			number, _ := strconv.Atoi(split3[1])
			foldings = append(foldings, Folding{direction: split3[0], line: number})
		}
	}
	for ii := 0; ii < folding && ii < len(foldings); ii++ {
		points = fold(points, foldings[ii])
	}
	return points
}

func fold(points []Point, fold Folding) []Point {
	pointSet := map[Point]bool{}
	if fold.direction == "y" {
		for _, dot := range points {
			if dot.y < fold.line {
				pointSet[dot] = true
			} else {
				dotNew := Point{y: 2*fold.line - dot.y, x: dot.x}
				pointSet[dotNew] = true
			}
		}
	} else if fold.direction == "x" {
		for _, dot := range points {
			if dot.x < fold.line {
				pointSet[dot] = true
			} else {
				dotNew := Point{y: dot.y, x: 2*fold.line - dot.x}
				pointSet[dotNew] = true
			}
		}
	}

	points = []Point{}
	for key := range pointSet {
		points = append(points, key)
	}
	return points
}

type Folding struct {
	direction string
	line      int
}

type Point struct {
	x, y int
}

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}

func Max(points []Point) (int, int) {
	var maxX int = points[0].x
	var maxY int = points[0].y
	for _, value := range points {
		if maxX < value.x {
			maxX = value.x
		}
		if maxY < value.y {
			maxY = value.y
		}
	}
	return maxX, maxY
}
