package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/beefsack/go-astar"
)

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}

type Element struct {
	x, y int
	maap *[][]int
}

func (e *Element) PathNeighbors() []astar.Pather {

	neighbors := []astar.Pather{}

	for _, yy := range []int{e.y + 1, e.y, e.y - 1} {
		for _, xx := range []int{e.x + 1, e.x, e.x - 1} {
			if numberNotOverEdge(*e.maap, yy, xx) && notDiagonal(e.x, e.y, xx, yy) {
				neighbors = append(neighbors, &Element{x: xx, y: yy, maap: e.maap})
			}
		}
	}
	return neighbors
}

// PathNeighborCost calculates the exact movement cost to neighbor nodes.
func (e *Element) PathNeighborCost(to astar.Pather) float64 {
	toE := to.(*Element)
	return float64((*(*e).maap)[toE.y][toE.x])
}

// PathEstimatedCost is a heuristic method for estimating movement costs
// between non-adjacent nodes.
func (e *Element) PathEstimatedCost(to astar.Pather) float64 {

	toT := to.(*Element)
	absX := toT.x - e.x
	if absX < 0 {
		absX = -absX
	}
	absY := toT.y - e.y
	if absY < 0 {
		absY = -absY
	}
	r := float64(absX + absY)

	return r
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

	maap := make([][]int, len(lines))
	for in, line := range lines {
		maap[in] = make([]int, len(line))
		for in2, run := range line {
			maap[in][in2] = int(run - '0')
		}
	}

	start := Element{x: 0, y: 0, maap: &maap}
	goal := Element{x: len(maap[0]), y: len(maap), maap: &maap}
	fmt.Println(start.PathNeighbors())
	path, zwei, ok := astar.Path(&start, &goal)

	if !ok {
		panic("this shit wont work")
	}
	fmt.Println(path, zwei)

	totalCosts := 0
	for _, elem := range path {
		element := elem.(*Element)
		totalCosts += element.getCosts()
	}
	fmt.Println(totalCosts)
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

func notDiagonal(xx, yy, xxNew, yyNew int) bool {
	return Abs(xx-xxNew)+Abs(yy-yyNew) == 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (e *Element) String() string {
	return fmt.Sprintf("Element[x:%v,y:%v]", e.x, e.y)
}

func (e *Element) getCosts() int {
	return (*(*e).maap)[e.y][e.x]
}
