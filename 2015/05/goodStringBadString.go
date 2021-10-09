package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var consonants *regexp.Regexp
var badStrings *regexp.Regexp

func main() {
	commandLineArgs := os.Args
	fileToRead := commandLineArgs[1]

	dat, _ := os.Open(fileToRead)

	scanner := bufio.NewScanner(dat)

	scannedStrings := []string{}
	for scanner.Scan() {
		scannedStrings = append(scannedStrings, scanner.Text())
	}

	CreateRegexp()
	niceStrings := 0
	for _, test := range scannedStrings {
		if IsStringGood(test) {
			niceStrings++
		}
	}

	fmt.Println("Number of nice strings first rule set:", niceStrings)

	niceStrings2 := 0
	for _, test := range scannedStrings {
		if IsStringGood2(test) {
			niceStrings2++
		}
	}

	fmt.Println("Number of nice strings second rule set:", niceStrings2)
}

func IsStringGood(test string) bool {

	good := true
	if !consonants.MatchString(test) {
		good = false
	}
	if badStrings.MatchString(test) {
		good = false
	}
	if !sameTwice(test) {
		good = false
	}
	return good
}

func IsStringGood2(test string) bool {

	good := true
	if !sameTwiceOneInBetween(test) {
		good = false
	}
	if !hateRule(test) {
		good = false
	}
	return good
}

func CreateRegexp() {
	consonants, _ = regexp.Compile("^.*[aeiou].*[aeiou].*[aeiou].*$")
	badStrings, _ = regexp.Compile(".*ab|cd|pq|xy.*$")
}

func sameTwice(test string) bool {
	var charBefore rune
	for _, char := range test {
		if char == charBefore {
			return true
		} else {
			charBefore = char
		}
	}
	return false
}

func sameTwiceOneInBetween(test string) bool {
	var charBeforeBefore rune
	var charBefore rune
	for _, char := range test {
		if char == charBeforeBefore {
			return true
		} else {
			charBeforeBefore = charBefore
			charBefore = char
		}
	}
	return false
}

func hateRule(test string) bool {
	for ii := 1; ii < len(test)-1; ii++ {
		if strings.Contains(test[ii+1:], test[ii-1:ii+1]) {
			return true
		}
	}
	return false
}
