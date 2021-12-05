package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type line struct {
	x1, y1, x2, y2 int
}

func main() {
	start := time.Now()
	dat, _ := os.Open("input.txt")
	defer func(dat *os.File) { dat.Close() }(dat)
	scanner := bufio.NewScanner(dat)
	inputLines := []line{}
	for scanner.Scan() {
		inString := scanner.Text()
		pairs := strings.Split(inString, "->")
		xy := []int{}
		for _, value := range pairs {
			numbers := strings.Split(value, ",")
			x, _ := strconv.Atoi(strings.Trim(numbers[0], " "))
			y, _ := strconv.Atoi(strings.Trim(numbers[1], " "))
			xy = append(xy, x)
			xy = append(xy, y)
		}
		inputLines = append(inputLines, line{x1: xy[0], x2: xy[2], y1: xy[1], y2: xy[3]})
	}

	max := 0
	for _, line := range inputLines {
		lineMax := line.findMaxValue()
		if lineMax > max {
			max = lineMax
		}
	}

	field := make([][]int, max+1)
	field2 := make([][]int, max+1)
	for ii := 0; ii < len(field); ii++ {
		field[ii] = make([]int, max+1)
		field2[ii] = make([]int, max+1)
	}

	for _, line := range inputLines {
		if line.isStraight() {
			line.sortLine()
			for xx := line.x1; xx <= line.x2; xx++ {
				for yy := line.y1; yy <= line.y2; yy++ {
					field[yy][xx]++
					field2[yy][xx]++
				}
			}
		} else {
			delta := int(math.Abs(float64(line.x2 - line.x1)))
			if delta != int(math.Abs(float64(line.y2-line.y1))) {
				panic("The Line is not diagonal")
			}
			if line.x1 > line.x2 {
			} else {
			}
			if line.y2 > line.y1 {
				for xx := 0; xx <= delta; xx++ {
					if line.x1 > line.x2 {
						field2[line.y1+xx][line.x1-xx]++
					} else {
						field2[line.y1+xx][line.x1+xx]++
					}
				}
			} else if line.y2 < line.y1 {
				for xx := 0; xx <= delta; xx++ {
					if line.x1 > line.x2 {
						field2[line.y1-xx][line.x1-xx]++
					} else {
						field2[line.y1-xx][line.x1+xx]++
					}
				}
			}
		}
	}

	overlappingPoints := 0
	for _, row := range field {
		for _, value := range row {
			if value > 1 {
				overlappingPoints++
			}
			//fmt.Printf("%v\t", value)
		}
		//fmt.Printf("\n")
	}
	fmt.Printf("The number of overlapping points for Part 1 is: %v\n", overlappingPoints)

	overlappingPoints2 := 0
	for _, row := range field2 {
		for _, value := range row {
			if value > 1 {
				overlappingPoints2++
			}
			//fmt.Printf("%v\t", value)
		}
		//fmt.Printf("\n")
	}
	fmt.Printf("The number of overlapping points for Part 2 is: %v\n", overlappingPoints2)
	elapsed := time.Since(start)
	log.Printf("This day took %s", elapsed)
}

func (l *line) findMaxValue() int {
	max := 0
	if l.x1 > max {
		max = l.x1
	}
	if l.x2 > max {
		max = l.x2
	}
	if l.y1 > max {
		max = l.y1
	}
	if l.y2 > max {
		max = l.y2
	}
	return max
}
func (l *line) sortLine() {
	dummy := 0
	if l.x1 > l.x2 {
		dummy = l.x1
		l.x1 = l.x2
		l.x2 = dummy
	}
	if l.y1 > l.y2 {
		dummy = l.y1
		l.y1 = l.y2
		l.y2 = dummy
	}
}
func (l *line) isStraight() bool {
	return l.x1 == l.x2 || l.y1 == l.y2
}
