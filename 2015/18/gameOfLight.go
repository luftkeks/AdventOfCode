package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dat, _ := os.Open("input.txt")
	defer func(dat *os.File) { dat.Close() }(dat)
	scanner := bufio.NewScanner(dat)
	lightField1 := []string{}
	lightField2 := []string{}
	for scanner.Scan() {
		lightField1 = append(lightField1, scanner.Text())
		lightField2 = append(lightField2, scanner.Text())
	}

	for ii := 0; ii < 100; ii++ {
		lightField1 = getNextIterationOfLightFiled(lightField1, false)
		lightField2 = getNextIterationOfLightFiled(lightField2, true)
	}

	counter1 := 0
	counter2 := 0
	for _, line := range lightField1 {
		for _, newChar := range line {
			if newChar == '#' {
				counter1++
			}
		}
	}
	for _, line := range lightField2 {
		for _, newChar := range line {
			if newChar == '#' {
				counter2++
			}
		}
	}

	// FUCK THIS SHIT - RESULT OF one IS Wrong!!!!
	fmt.Printf("Number of Lights on is %v .\n", counter1)
	fmt.Printf("Number of Lights which are on in the second task is %v .", counter2)
}

func getNextIterationOfLightFiled(lightField []string, secondTask bool) (lightFieldFuture []string) {
	lightFieldFuture = make([]string, len(lightField))
	for line, row := range lightField {
		lightFieldFuture[line] = row
		for symbol, char := range lightField[line] {
			counter := getCounterForChar(lightField, line, symbol)
			var newChar rune
			if char == '#' {
				if counter == 2 || counter == 3 {
					newChar = '#'
				} else {
					newChar = '.'
				}
			} else if char == '.' {
				if counter == 3 {
					newChar = '#'
				} else {
					newChar = '.'
				}
			}
			if secondTask && (line == 0 && symbol == 0) || (line == 0 && symbol == len(lightField[0])-1) || (line == len(lightField)-1 && symbol == 0) || (line == len(lightField)-1 && symbol == len(lightField[0])-1) {
				newChar = '#'
			}
			row := lightFieldFuture[line]
			lightFieldFuture[line] = fmt.Sprintf("%v%v%v", row[:symbol], string(newChar), row[symbol+1:])
		}
	}
	return
}

func getCounterForChar(lightField []string, line, symbol int) int {
	counter := 0
	if line > 0 && symbol > 0 {
		if lightField[line-1][symbol-1] == '#' {
			counter++
		}
	}
	if line > 0 {
		if lightField[line-1][symbol] == '#' {
			counter++
		}
	}
	if line > 0 && symbol < len(lightField[0])-1 {
		if lightField[line-1][symbol+1] == '#' {
			counter++
		}
	}
	if symbol > 0 {
		if lightField[line][symbol-1] == '#' {
			counter++
		}
	}
	if symbol < len(lightField[0])-1 {
		if lightField[line][symbol+1] == '#' {
			counter++
		}
	}
	if line < len(lightField)-1 && symbol > 0 {
		if lightField[line+1][symbol-1] == '#' {
			counter++
		}
	}
	if line < len(lightField)-1 {
		if lightField[line+1][symbol] == '#' {
			counter++
		}
	}
	if line < len(lightField)-1 && symbol < len(lightField[0])-1 {
		if lightField[line+1][symbol+1] == '#' {
			counter++
		}
	}
	return counter
}
