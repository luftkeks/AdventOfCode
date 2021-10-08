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

	for _, char := range string(dat) {
		if char == '(' {
			floor++
		}
		if char == ')' {
			floor--
		}
	}

	fmt.Println(floor)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
