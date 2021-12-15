package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"time"
)

type MapElement struct {
	x, y int
	cost int
	maap *[][]*MapElement
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

	maap := make([][]*MapElement, len(lines))
	for in, line := range lines {
		maap[in] = make([]*MapElement, len(line))
		for in2, run := range line {
			maap[in][in2] = &MapElement{x: in2, y: in, cost: int(run - '0'), maap: &maap}
		}
	}

	start := maap[0][0]
	goal := maap[len(maap)-1][len(maap[0])-1]
	path := aStarAlgorithm(start, goal)

	totalCosts := 0
	for _, elem := range path {
		totalCosts += elem.cost
	}
	fmt.Printf("The total cost for part 1 is: %v\n", totalCosts)

	maap2 := make([][]*MapElement, len(lines)*5)
	for ii := 0; ii < len(lines)*5; ii++ {
		line := lines[ii%len(lines)]
		maap2[ii] = make([]*MapElement, len(line)*5)
		for jj := 0; jj < len(line)*5; jj++ {
			maap2[ii][jj] = &MapElement{x: jj, y: ii, cost: wrapNine(int(rune(line[jj%len(line)])-'0') + ii/len(line) + jj/len(lines)), maap: &maap2}
		}
	}

	start2 := maap2[0][0]
	goal2 := maap2[len(maap2)-1][len(maap2[0])-1]
	path2 := aStarAlgorithm(start2, goal2)
	totalCosts2 := 0
	for _, elem := range path2 {
		totalCosts2 += elem.cost
	}
	fmt.Printf("The total cost for part 2 is: %v\n", totalCosts2)
}

func wrapNine(number int) int {
	if number > 9 {
		return wrapNine(number - 9)
	} else {
		return number
	}
}

// this code ist copied from https://github.com/feliperyan/golang-astar-algo and modified by me
func aStarAlgorithm(start *MapElement, goal *MapElement) []*MapElement {

	frontier := make(PriorityQueue, 0)
	heap.Init(&frontier)
	heap.Push(&frontier, &Item{value: start, priority: 0})

	cameFrom := make(map[*MapElement]*MapElement)
	costSoFar := make(map[*MapElement]int)
	cameFrom[start] = nil
	costSoFar[start] = 0

	finished := false

	for i := len(frontier); i > 0; {

		topItem := heap.Pop(&frontier).(*Item)
		current := topItem.value

		if current.x == goal.x && current.y == goal.y {
			finished = true
			break
		}

		neighbours := current.possibleNeighbor()

		for _, nextElement := range neighbours {
			newCost := costSoFar[current] + nextElement.cost
			_, exists := costSoFar[nextElement]

			if (!exists) || (newCost < costSoFar[nextElement]) {
				costSoFar[nextElement] = newCost
				priority := newCost + int(goal.getEstimatedCost(nextElement))
				heap.Push(&frontier, &Item{value: nextElement, priority: priority})
				cameFrom[nextElement] = current
			}
		}

		i = len(frontier)
	}

	if finished {
		path := getPathToGoal(cameFrom, start, goal)
		return path
	} else {
		path := make([]*MapElement, 0)
		return path
	}
}

func (e *MapElement) possibleNeighbor() []*MapElement {

	neighbors := []*MapElement{}

	for _, yy := range []int{e.y + 1, e.y, e.y - 1} {
		for _, xx := range []int{e.x + 1, e.x, e.x - 1} {
			if numberNotOverEdge(*e.maap, yy, xx) && notDiagonal(e.x, e.y, xx, yy) {
				neighbors = append(neighbors, (*(*e).maap)[yy][xx])
			}
		}
	}
	return neighbors
}

func (e *MapElement) getEstimatedCost(to *MapElement) float64 {

	absX := to.x - e.x
	if absX < 0 {
		absX = -absX
	}
	absY := to.y - e.y
	if absY < 0 {
		absY = -absY
	}
	r := float64(absX + absY)

	return r
}

func getPathToGoal(cameFrom map[*MapElement]*MapElement, start, goal *MapElement) []*MapElement {
	nodes := make([]*MapElement, 0)

	for e := goal; e != start; {
		// prepend trick so I don't have to reverse the list later
		nodes = append([]*MapElement{e}, nodes...) // this is prob inneficient.
		e = cameFrom[e]
	}
	return nodes
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    *MapElement // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.

}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// ATTENTION: For the A* Algorithm I want the item with the LOWEST priority (cost)
	// so the sign is inverted from the original code and regular priority queues.
	return pq[i].priority <= pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value *MapElement, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func numberNotOverEdge(maap [][]*MapElement, yy, xx int) bool {
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

func (e *MapElement) String() string {
	return fmt.Sprintf("Element[x:%v,y:%v]", e.x, e.y)
}
