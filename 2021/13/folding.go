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

	dots := parseWithFoldings("input.txt", 1)
	fmt.Println(dots)
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
	return len(fold(points, foldings[0]))
}

func fold(points []Point, fold Folding) []Point {
	pointSet := map[Point]bool{}
	if fold.direction == "y" {

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
