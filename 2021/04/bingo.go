package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoField struct {
	field [][]int
}

func main() {
	dat, _ := os.Open("input.txt")
	defer func(dat *os.File) { dat.Close() }(dat)
	scanner := bufio.NewScanner(dat)
	inputLines := []string{}
	for scanner.Scan() {
		inString := scanner.Text()
		inputLines = append(inputLines, inString)
	}

	input := strings.Split(inputLines[0], ",")

	checkForPart1(initBingos(inputLines), input)

	bingos2 := initBingos(inputLines)
	for _, number := range input {
		value, _ := strconv.Atoi(number)

		itemsToRemove := map[int]bool{}
		for ii, bingo := range bingos2 {
			if bingo.checkForNumber(value) {
				if bingo.checkForBingo() {
					if len(itemsToRemove)+1 != len(bingos2) {
						itemsToRemove[ii] = true
					} else {
						fmt.Printf("The Result 2 should be: %v, last value was: %v \n", value*bingos2[0].addRemainingNumbers(), value)
						fmt.Println(bingos2[0].String())
						return
					}
				}
			}
		}
		if len(itemsToRemove) > 0 {
			dummy := []BingoField{}
			for jj, val := range bingos2 {
				ok := itemsToRemove[jj]
				if !ok {
					dummy = append(dummy, val)
				}
			}
			bingos2 = dummy
		}
	}

}

func checkForPart1(bingos []BingoField, input []string) {
	for _, number := range input {
		value, _ := strconv.Atoi(number)

		for _, bingo := range bingos {
			if bingo.checkForNumber(value) {
				if bingo.checkForBingo() {
					fmt.Printf("The Result 1 should be: %v, last value was: %v \n", value*bingo.addRemainingNumbers(), value)
					fmt.Println(bingo.String())
					return
				}
			}
		}
	}
}

func (b *BingoField) checkForNumber(number int) bool {
	for ll, line := range b.field {
		for rr, value := range line {
			if value == number {
				b.field[ll][rr] = 0
				return true
			}
		}
	}
	return false
}

func (b *BingoField) checkForBingo() bool {
	for _, line := range b.field {
		if checkLine(line) {
			return true
		}
	}

	for ii := 0; ii < 5; ii++ {
		thing := []int{}
		for _, line := range b.field {
			thing = append(thing, line[ii])
		}
		if checkLine(thing) {
			return true
		}
	}

	return false
}

func (b *BingoField) addRemainingNumbers() (result int) {
	for _, xx := range b.field {
		for _, yy := range xx {
			result += yy
		}
	}
	return
}

func checkLine(line []int) bool {
	for _, value := range line {
		if value != 0 {
			return false
		}
	}
	return true
}

func (b *BingoField) String() string {
	var result string
	for _, line := range b.field {
		for _, value := range line {
			result += strconv.Itoa(value) + " "
		}
		result += "\n"
	}
	return result
}

func initBingos(inputLines []string) []BingoField {
	bingos := []BingoField{{field: [][]int{}}}

	for ii, line := range inputLines[2:] {
		if (ii+1)%6 == 0 {
			bingos = append(bingos, BingoField{field: [][]int{}})
			continue
		}
		values := strings.Split(line, " ")
		row := []int{}
		for _, value := range values {
			number, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			row = append(row, number)
		}
		bingos[len(bingos)-1].field = append(bingos[len(bingos)-1].field, row)
	}
	return bingos
}
