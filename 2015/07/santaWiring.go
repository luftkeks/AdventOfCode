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

type Wiring struct {
	Wiring map[string]Wire
}

func (v Wire) GetUint(wiring map[string]Wire) uint16 {
	if v.Solved {
		return v.value
	} else if v.operation == "NOT" {
		v.value = ^wiring[v.ref1].GetUint(wiring)
		v.Solved = true
	} else if v.operation == "IS" {
		v.value = wiring[v.ref1].GetUint(wiring)
		v.Solved = true
	} else if v.operation == "AND" {
		v.value = wiring[v.ref1].GetUint(wiring) & wiring[v.ref2].GetUint(wiring)
		v.Solved = true
	} else if v.operation == "OR" {
		v.value = wiring[v.ref1].GetUint(wiring) | wiring[v.ref2].GetUint(wiring)
		v.Solved = true
	} else if v.operation == "LSHIFT" {
		v.value = wiring[v.ref1].GetUint(wiring) << v.shift
		v.Solved = true
	} else if v.operation == "RSHIFT" {
		v.value = wiring[v.ref1].GetUint(wiring) >> v.shift
		v.Solved = true
	}
	fmt.Println(v.Name)
	return v.value
}

func (w *Wiring) Get(name string) Wire {
	return w.Wiring[name]
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

	wiring := map[string]Wire{}

	WireMap(scannedStrings, wiring)

	fmt.Println("The Value of a is: ", wiring["a"].GetUint(wiring))
}

func WireMap(stringList []string, wiring map[string]Wire) {
	for _, line := range stringList {
		wire := parseLine(line)
		wiring[wire.Name] = wire
	}
}

func parseLine(line string) (node Wire) {
	words := strings.Split(line, " ")
	if len(words) == 3 {
		if isStringNumber(words[0]) {
			number, _ := strconv.Atoi(words[0])
			return Wire{Solved: true, value: uint16(number), Name: words[2]}
		} else {
			return Wire{ref1: words[0], operation: "IS", Name: words[2], Solved: false}
		}
	} else if len(words) == 4 && words[0] == "NOT" {
		return Wire{ref1: words[1], operation: "NOT", Name: words[3], Solved: false}
	} else if len(words) == 5 && (words[1] == "AND" || words[1] == "OR") {
		return Wire{ref1: words[0], ref2: words[2], Name: words[4], Solved: false}
	} else if len(words) == 5 && (words[1] == "LSHIFT" || words[1] == "RSHIFT") {
		numb, _ := strconv.Atoi(words[2])
		return Wire{ref1: words[0], shift: uint16(numb), Name: words[4], Solved: false}
	}
	panic("HELP")
}

func isStringNumber(word string) bool {
	_, err := strconv.Atoi(word)
	if err == nil {
		return true
	} else {
		return false
	}
}
