package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Wire struct {
	Name      string
	value     uint16
	Solved    bool
	ref1      string
	ref2      string
	shift     uint16
	operation string
}

func (w *Wire) GetUint(wiring map[string]*Wire, caller string) uint16 {
	if w == nil {
		panic("FUCKING HELL " + caller + " What did you do?")
	}
	if w.Solved {
		return w.value
	}

	value1 := valueForWire(w.ref1, wiring, caller)

	if w.operation == "NOT" {
		w.value = ^value1
		w.Solved = true
	} else if w.operation == "IS" {
		w.value = value1
		w.Solved = true
	} else if w.operation == "LSHIFT" {
		w.value = value1 << w.shift
		w.Solved = true
	} else if w.operation == "RSHIFT" {
		w.value = value1 >> w.shift
		w.Solved = true
	} else if w.operation == "AND" {
		w.value = value1 & valueForWire(w.ref2, wiring, caller)
		w.Solved = true
	} else if w.operation == "OR" {
		w.value = value1 | valueForWire(w.ref2, wiring, caller)
		w.Solved = true
	}
	return w.value
}

func valueForWire(wireName string, wiring map[string]*Wire, caller string) (value uint16) {
	number, err := strconv.Atoi(wireName)
	if err != nil {
		wire := wiring[wireName]
		value = wire.GetUint(wiring, caller)
	} else {
		value = uint16(number)
	}
	return value
}

func main() {
	commandLineArgs := os.Args
	fileToRead := commandLineArgs[1]
	dat, _ := os.Open(fileToRead)
	scanner := bufio.NewScanner(dat)
	scannedStrings := []string{}
	for scanner.Scan() {
		scannedStrings = append(scannedStrings, scanner.Text())
	}

	wiring := map[string]*Wire{}

	WireMap(scannedStrings, wiring)

	wireA := wiring["a"]
	signalA := wireA.GetUint(wiring, "main class")
	fmt.Println("The Value of a is: ", signalA)

	wiring2 := map[string]*Wire{}
	WireMap(scannedStrings, wiring2)
	wiring2["b"] = &Wire{value: signalA, operation: "IS", Name: "b", Solved: true}

	wireA2 := wiring2["a"]
	signalA2 := wireA2.GetUint(wiring2, "main class")
	fmt.Println("The Value of a is after the change: ", signalA2)
}

func WireMap(stringList []string, wiring map[string]*Wire) {
	for _, line := range stringList {
		wire := parseLine(line)
		wiring[wire.Name] = &wire
	}
}

func parseLine(line string) (node Wire) {
	words := strings.Split(line, " ")
	if len(words) == 3 {
		number, err := strconv.Atoi(words[0])
		if err == nil {
			return Wire{value: uint16(number), operation: "IS", Name: words[2], Solved: true}
		} else {
			return Wire{ref1: words[0], operation: "IS", Name: words[2], Solved: false}
		}
	} else if len(words) == 4 && words[0] == "NOT" {
		return Wire{ref1: words[1], operation: words[0], Name: words[3], Solved: false}
	} else if len(words) == 5 && (words[1] == "AND" || words[1] == "OR") {
		return Wire{ref1: words[0], ref2: words[2], operation: words[1], Name: words[4], Solved: false}
	} else if len(words) == 5 && (words[1] == "LSHIFT" || words[1] == "RSHIFT") {
		numb, _ := strconv.Atoi(words[2])
		return Wire{ref1: words[0], shift: uint16(numb), operation: words[1], Name: words[4], Solved: false}
	}
	panic("HELP")
}
