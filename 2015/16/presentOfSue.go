package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.Open("input.txt")
	defer func(dat *os.File) { dat.Close() }(dat)
	scanner := bufio.NewScanner(dat)
	scannedStrings := []string{}
	for scanner.Scan() {
		scannedStrings = append(scannedStrings, strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), ":", ""), ",", ""))
	}

	remeberedSues := []map[string]int{}

	for _, line := range scannedStrings {
		words := strings.Split(line, " ")
		sue := map[string]int{}
		first, err := strconv.Atoi(words[3])
		if err != nil {
			log.Fatalln(err)
		}
		second, err := strconv.Atoi(words[5])
		if err != nil {
			log.Fatalln(err)
		}
		third, err := strconv.Atoi(words[7])
		if err != nil {
			log.Fatalln(err)
		}
		sue[words[2]] = first
		sue[words[4]] = second
		sue[words[6]] = third
		remeberedSues = append(remeberedSues, sue)
	}

	rightSue := map[string]int{"children": 3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1}

	for index, testSue := range remeberedSues {
		test := true
		for key, value := range testSue {
			if rightSue[key] != value {
				test = false
			}
		}
		if test {
			fmt.Printf("Sue number %v is the right one before correction!\n", index+1)
		}
	}

	for index, testSue := range remeberedSues {
		test := true
		for key, value := range testSue {
			if key == "cats" || key == "trees" {
				if rightSue[key] > value {
					test = false
				}
			} else if key == "pomeranians" || key == "goldfish" {
				if rightSue[key] < value {
					test = false
				}
			} else if rightSue[key] != value {
				test = false
			}
		}
		if test {
			fmt.Printf("Sue number %v is the right one after correction!\n", index+1)
		}
	}
}
