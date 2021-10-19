package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	dat, _ := os.Open("input.txt")
	defer func(dat *os.File) { dat.Close() }(dat)
	scanner := bufio.NewScanner(dat)
	scannedNumbers := []int{}
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		scannedNumbers = append(scannedNumbers, number)
	}

	lenNumberMap := map[int]int{}
	numberOfSolutions := subset_sum(scannedNumbers, 150, []int{}, &lenNumberMap)

	fmt.Printf("There are %v possible solutions.\n", numberOfSolutions)

	fmt.Println(lenNumberMap)
}

// borrowed from https://stackoverflow.com/questions/4632322/finding-all-possible-combinations-of-numbers-to-reach-a-given-sum
func subset_sum(numbers []int, target int, partial []int, lenNumberMap *map[int]int) int {
	s := sum(partial)

	// check if the partial sum is equals to target
	if s == target {
		fmt.Printf("%v adds up to: %v\n", partial, target)
		(*lenNumberMap)[len(partial)] += 1
		return 1
	}
	if s >= target {
		return 0
	} // if we reach the number why bother to continue
	numberOfSolutions := 0
	for i := range numbers {
		n := numbers[i]
		remaining := numbers[i+1:]
		numberOfSolutions += subset_sum(remaining, target, append(partial, n), lenNumberMap)
	}
	return numberOfSolutions
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
