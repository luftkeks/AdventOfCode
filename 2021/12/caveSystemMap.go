package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}

func main() {
	defer elapsed()()

	nodes := loadNodes("input.txt")
	fmt.Printf("The Number of valid Paths for Part 1 is: %v\n", Part1(nodes))
	fmt.Printf("The Number of valid Paths for Part 1 is: %v\n", Part2(nodes))

}

func Part1(nodes map[string]*Node) (out int) {

	startNode := nodes["start"]

	ways := findWay(startNode, nodes, Way{startNode}, 20, false)

	counter := 0
	for _, way := range ways {
		if way.isValid1() && way.hasEnded() {
			counter++
		}
	}
	return counter
}

func Part2(nodes map[string]*Node) (out int) {

	startNode := nodes["start"]

	ways := findWay(startNode, nodes, Way{startNode}, 20, true)

	counter := 0
	for _, way := range ways {
		if way.isValid2() && way.hasEnded() {
			counter++
		}
	}
	return counter
}

type Node struct {
	name   string
	edgeds map[string]bool
}

type Way []*Node

func findWay(node *Node, nodes map[string]*Node, way Way, depth int, isPart2 bool) []Way {
	if depth < 0 {
		return []Way{}
	}
	if node.name == "end" {
		return []Way{Way{node}}
	}

	returnWays := []Way{}
	for newNode := range node.edgeds {
		nNode := nodes[newNode]
		newWay := append(way, nNode)
		if (!isPart2 && newWay.isValid1()) || (isPart2 && newWay.isValid2()) {
			ways := findWay(nNode, nodes, newWay, depth-1, isPart2)
			returnWays = append(returnWays, ways...)
		}
	}
	for ii := 0; ii < len(returnWays); ii++ {
		returnWays[ii] = append(returnWays[ii], node)
	}
	return returnWays
}

func (n *Node) addEdge(in string) {
	if n.edgeds == nil {
		n.edgeds = make(map[string]bool)
	}
	n.edgeds[in] = true
}

func (n *Node) String() string {
	return fmt.Sprintf("[%v:%v]", n.name, n.edgeds)
}

func (w *Way) String() string {
	result := ""
	for ii := len(*w) - 1; ii >= 0; ii-- {
		result = result + " " + (*w)[ii].name
	}
	return result
}

func (w *Way) isValid1() bool {
	nodeNumberMap := make(map[string]int)

	for ii := 0; ii < len(*w); ii++ {
		nodeNumberMap[(*w)[ii].name]++
	}

	for kk, vv := range nodeNumberMap {
		if rune(kk[0]) > 'a' && vv > 1 {
			return false
		}
	}
	return true
}

func (w *Way) isValid2() bool {
	nodeNumberMap := make(map[string]int)

	for ii := 0; ii < len(*w); ii++ {
		nodeNumberMap[(*w)[ii].name]++
	}

	once := true
	for kk, vv := range nodeNumberMap {
		if kk != "start" && rune(kk[0]) >= 'a' && vv == 2 && once {
			once = false
			continue
		}
		if rune(kk[0]) >= 'a' && vv > 1 {
			return false
		}
	}
	return true
}

func (w *Way) hasEnded() bool {
	if (*w)[0].name == "end" {
		return true
	} else {
		return false
	}
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
