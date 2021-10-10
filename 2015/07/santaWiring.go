package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	x string
}

type And struct {
	x, y string
}

type Or struct {
	x, y string
}

type Not struct {
	x string
}

type Rshift struct {
	x      string
	number int
}

type Lshift struct {
	x      string
	number int
}

type Dummy struct {
	number uint16
}

type Wire interface {
	Get(Wiring) uint16
}

type Wiring map[string]Wire

func (w Wiring) Get(name string) Wire {
	if isStringNumber(name) {
		numb, _ := strconv.Atoi(name)
		return Dummy{number: uint16(numb)}
	} else {
		return w[name]
	}
}

func (d Dummy) Get(wiring Wiring) uint16 {
	return d.number
}

func (s Set) Get(wiring Wiring) uint16 {
	return wiring.Get(s.x).Get(wiring)
}
func (a And) Get(wiring Wiring) uint16 {
	return wiring.Get(a.x).Get(wiring) & wiring.Get(a.y).Get(wiring)
}
func (o Or) Get(wiring Wiring) uint16 {
	return wiring.Get(o.x).Get(wiring) | wiring.Get(o.y).Get(wiring)
}
func (l Lshift) Get(wiring Wiring) uint16 {
	return wiring.Get(l.x).Get(wiring) << uint16(l.number)
}
func (r Rshift) Get(wiring Wiring) uint16 {
	return wiring.Get(r.x).Get(wiring) >> uint16(r.number)
}
func (n Not) Get(wiring Wiring) uint16 { return ^wiring.Get(n.x).Get(wiring) }

func main() {
	commandLineArgs := os.Args
	fileToRead := commandLineArgs[1]
	dat, _ := os.Open(fileToRead)
	scanner := bufio.NewScanner(dat)
	scannedStrings := []string{}
	for scanner.Scan() {
		scannedStrings = append(scannedStrings, scanner.Text())
	}

	wiring := Wiring{}

	WireMap(scannedStrings, wiring)

	fmt.Println("The Value of a is: ", wiring.Get("a").Get(wiring))
}

func WireMap(stringList []string, wiring map[string]Wire) {
	for _, line := range stringList {
		wire, name := parseLine(line, wiring)
		wiring[name] = wire
	}
}

func parseLine(line string, wiring map[string]Wire) (node Wire, name string) {
	words := strings.Split(line, " ")
	if len(words) == 3 {
		return Set{x: words[0]}, words[2]
	} else if len(words) == 4 {
		return Not{x: words[1]}, words[3]
	} else if len(words) == 5 && words[1] == "AND" {
		return And{x: words[0], y: words[2]}, words[4]
	} else if len(words) == 5 && words[1] == "OR" {
		return Or{x: words[0], y: words[2]}, words[4]
	} else if len(words) == 5 && words[1] == "LSHIFT" {
		numb, _ := strconv.Atoi(words[2])
		return Lshift{x: words[0], number: numb}, words[4]
	} else if len(words) == 5 && words[1] == "RSHIFT" {
		numb, _ := strconv.Atoi(words[2])
		return Rshift{x: words[0], number: numb}, words[4]
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
