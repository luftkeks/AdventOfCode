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
	cost int
	maap *[][]Element
}

func (e *Element) PathNeighbors() []astar.Pather {

	neighbors := []astar.Pather{}

	for _, yy := range []int{e.y + 1, e.y, e.y - 1} {
		for _, xx := range []int{e.x + 1, e.x, e.x - 1} {
			if numberNotOverEdge(*e.maap, yy, xx) && notDiagonal(e.x, e.y, xx, yy) {
				neighbors = append(neighbors, &(*(*e).maap)[yy][xx])
			}
		}
	}
	return neighbors
}

// PathNeighborCost calculates the exact movement cost to neighbor nodes.
func (e *Element) PathNeighborCost(to astar.Pather) float64 {
	toE := to.(*Element)
	return float64((*(*e).maap)[toE.y][toE.x].cost)
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

	maap := make([][]Element, len(lines))
	for in, line := range lines {
		maap[in] = make([]Element, len(line))
		for in2, run := range line {
			maap[in][in2] = Element{x: in2, y: in, cost: int(run - '0'), maap: &maap}
		}
	}

	start := Element{x: 0, y: 0, maap: &maap}
	goal := Element{x: len(maap[0]) - 1, y: len(maap) - 1, maap: &maap}
	for _, elem := range start.PathNeighbors() {
		fmt.Println(elem.PathNeighborCost(&start))
	}

	fmt.Println(goal.PathNeighbors())
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
	fmt.Printf("The total cost minus the start element are: %v\n", totalCosts-maap[0][0].cost)
}

func numberNotOverEdge(maap [][]Element, yy, xx int) bool {
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
	return e.cost
}
