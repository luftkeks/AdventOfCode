package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	name     string
	relation map[string]int
}

func main() {
	dat, _ := os.Open("input.txt")
	defer func(dat *os.File) { dat.Close() }(dat)
	scanner := bufio.NewScanner(dat)
	scannedStrings := []string{}
	for scanner.Scan() {
		scannedStrings = append(scannedStrings, scanner.Text())
	}

	persons := map[string]Person{}
	for _, line := range scannedStrings {
		persons = parseInput(line, persons)
	}

	fmt.Println("Biggest Happines without me is:", findBiggestHappiness(persons))
	keys := make([]string, 0, len(persons))
	for k := range persons {
		keys = append(keys, k)
	}
	me := Person{name: "Me", relation: map[string]int{}}
	for _, person := range persons {
		person.relation["Me"] = 0
	}
	for _, name := range keys {
		me.relation[name] = 0
	}
	persons["me"] = me
	fmt.Println("Biggest Happines with me is:", findBiggestHappiness(persons))
}

func findBiggestHappiness(persons map[string]Person) int {
	keys := make([]string, 0, len(persons))
	for k := range persons {
		keys = append(keys, k)
	}
	permut := permutations(keys)

	channel := make(chan int)

	numberOfPermut := 0
	biggestHappines := 0
	for _, order := range permut {
		go checkOrder(order, persons, channel)
		numberOfPermut += 1
	}

	for ii := 0; ii < numberOfPermut; ii++ {
		thing := <-channel
		if thing > biggestHappines {
			biggestHappines = thing
		}
	}
	return biggestHappines
}

func parseInput(input string, persons map[string]Person) map[string]Person {
	words := strings.Split(input, " ")
	name := words[0]
	number, _ := strconv.Atoi(words[3])
	nextTo := words[10][:len(words[10])-1]

	if words[2] == "lose" {
		number *= -1
	}

	person, ok := persons[name]
	if ok {
		person.mergePerson(Person{name: name, relation: map[string]int{nextTo: number}})
	} else {
		persons[name] = Person{name: name, relation: map[string]int{nextTo: number}}
	}
	return persons
}

func (p *Person) mergePerson(other Person) {
	if p.name == other.name {
		for name, number := range other.relation {
			p.relation[name] = number
		}
	}
}

func checkOrder(list []string, persons map[string]Person, c chan int) {
	round := func(number int, list []string) int {
		if number >= len(list) {
			return number - len(list)
		} else if number < 0 {
			return number + len(persons)
		}
		return number
	}
	result := 0
	for index, name := range list {
		person := persons[name]
		after := list[round(index+1, list)]
		before := list[round(index-1, list)]
		result += person.relation[after]
		result += person.relation[before]
	}
	c <- result
}

// Not my function - copied from: https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
