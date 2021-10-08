package main

import (
	"fmt"
	"os"
)

func main() {
	commandLineArgs := os.Args
	fileToRead := commandLineArgs[1]

	dat, err := os.ReadFile(fileToRead)
	check(err)

	floor := 0
	var position int

	for number, char := range string(dat) {
		if char == '(' {
			floor++
		}
		if char == ')' {
			floor--
		}
		if floor == -1 && position == 0 {
			position = number + 1
		}
	}

	fmt.Printf("Position where he ends: ")
	fmt.Println(floor)

	fmt.Printf("Position where he first enters the basement: ")
	fmt.Println(position)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
