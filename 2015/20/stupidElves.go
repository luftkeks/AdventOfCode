package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	commandLineArgs := os.Args
	magicNumber, err := strconv.Atoi(commandLineArgs[1])
	if err != nil {
		panic("What have you done, you f*cking idiot!")
	}

	numbers := []int{0, 0}

	counter := 1
	for {
		gifts := 0
		numbers = append(numbers, 0)
		for ii := 1; ii <= counter; ii++ {
			if counter%ii == 0 && numbers[ii] <= 50 {
				numbers[ii] += 1
				gifts += ii * 11
			}
		}
		if gifts >= magicNumber {
			break
		}
		counter++
	}

	fmt.Printf("The first house with %v gifts is: %v !\n", magicNumber, counter)
	fmt.Println("Andi is a oarsch mit oarn!")
}
