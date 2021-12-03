package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dat, _ := os.Open("input.txt")
	defer func(dat *os.File) { dat.Close() }(dat)
	scanner := bufio.NewScanner(dat)
	strings := []string{}
	for scanner.Scan() {
		inString := scanner.Text()
		strings = append(strings, inString)
	}

	resultSlice := make([]uint, len(strings[0]))
	for _, number := range strings {
		for ii, thing := range number {
			if thing == '1' {
				resultSlice[ii]++
			}
		}
	}

	var result1 string
	var result2 string
	for _, number := range resultSlice {
		if number > uint(len(strings))/2 {
			result1 += "1"
			result2 += "0"
		} else {
			result1 += "0"
			result2 += "1"

		}
	}

	res1Most, _ := strconv.ParseInt(result1, 2, 0)
	res1Least, _ := strconv.ParseInt(result2, 2, 0)
	fmt.Printf("The Result of 1 is: %v\n", res1Most*res1Least)

	result21 := findMost(strings, 0)
	result22 := findLeast(strings, 0)
	res2Most, _ := strconv.ParseInt(result21[0], 2, 0)
	res2Least, _ := strconv.ParseInt(result22[0], 2, 0)
	fmt.Printf("The Numbers for Result 2 are: %v and %v\n", result21[0], result22[0])
	fmt.Printf("The Result of 2 is: %v\n", res2Most*res2Least)
}

func findMost(strings []string, location int) []string {
	numberOfOne := 0
	for _, str := range strings {
		if str[location] == '1' {
			numberOfOne++
		}
	}

	var test rune
	if float64(numberOfOne) >= float64(len(strings))/2.0 {
		test = '1'
	} else {
		test = '0'
	}

	result := []string{}
	for _, str := range strings {
		if str[location] == byte(test) {
			result = append(result, str)
		}
	}
	if len(result) == 1 {
		return result
	} else {
		return findMost(result, location+1)
	}
}

func findLeast(strings []string, location int) []string {
	numberOfOne := 0
	for _, str := range strings {
		if str[location] == '1' {
			numberOfOne++
		}
	}

	var test rune
	if float64(numberOfOne) >= float64(len(strings))/2.0 {
		test = '0'
	} else {
		test = '1'
	}

	result := []string{}
	for _, str := range strings {
		if str[location] == byte(test) {
			result = append(result, str)
		}
	}
	if len(result) == 1 {
		return result
	} else {
		return findLeast(result, location+1)
	}
}
