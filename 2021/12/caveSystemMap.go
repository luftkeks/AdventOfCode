package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type count32 int32
type Node struct {
	name  string
	edges map[string]bool
}

type Way []*Node

func main() {
	defer elapsed()()

	nodes := loadNodes("input.txt")
	part1, part2 := Part(nodes)
	fmt.Printf("The Number of valid Paths for Part 1 is: %v\n", part1)
	fmt.Printf("The Number of valid Paths for Part 1 is: %v\n", part2)

}

func Part(nodes map[string]*Node) (out1, out2 int) {

	startNode := nodes["start"]

	wait := sync.WaitGroup{}
	wait.Add(1)

	counter1 := count32(0)
	counter2 := count32(0)
	findWay(nodes, Way{startNode}, &counter1, &counter2, &wait)

	wait.Wait()
	return int(counter1.get()), int(counter2.get())
}

func findWay(nodes map[string]*Node, way Way, count1 *count32, count2 *count32, wait *sync.WaitGroup) {
	node := way[0]
	if node.name == "end" && way.isValid(true) {
		if way.isValid(false) {
			count1.inc()
		}
		count2.inc()
		wait.Done()
		return
	} else if !way.isValid(true) {
		wait.Done()
		return
	}

	wait.Add(len(node.edges) - 1)
	for newNode := range node.edges {
		newWay := append(Way{nodes[newNode]}, way...)
		go findWay(nodes, newWay, count1, count2, wait)
	}
}

func (n *Node) addEdge(in string) {
	if n.edges == nil {
		n.edges = make(map[string]bool)
	}
	n.edges[in] = true
}

func (w *Way) isValid(isPartTwo bool) bool {
	nodeNumberMap := make(map[string]int)

	for ii := 0; ii < len(*w); ii++ {
		nodeNumberMap[(*w)[ii].name]++
	}

	for kk, vv := range nodeNumberMap {
		if isPartTwo && kk != "start" && rune(kk[0]) >= 'a' && vv == 2 {
			isPartTwo = false
			continue
		}
		if rune(kk[0]) >= 'a' && vv > 1 {
			return false
		}
	}
	return true
}

func loadNodes(inputFile string) map[string]*Node {
	dat, err := os.Open(inputFile)
	if err != nil {
		panic("Hilfe File tut nicht")
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var nodes = map[string]*Node{}
	for scanner.Scan() {
		inString := scanner.Text()
		node01 := strings.Split(inString, "-")
		Pnode1 := nodes[node01[0]]
		var node1 Node
		if Pnode1 != nil {
			node1 = *Pnode1
		} else {
			node1 = Node{name: node01[0]}
		}
		node1.addEdge(node01[1])
		nodes[node01[0]] = &node1

		Pnode2 := nodes[node01[1]]
		var node2 Node
		if Pnode2 != nil {
			node2 = *Pnode2
		} else {
			node2 = Node{name: node01[1]}
		}
		node2.addEdge(node01[0])
		nodes[node01[1]] = &node2
	}
	return nodes
}

func (c *count32) inc() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *count32) get() int32 {
	return atomic.LoadInt32((*int32)(c))
}

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}
