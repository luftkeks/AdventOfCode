package main

import (
	"fmt"
)

func main() {
	test := "1321131112"

	for ii := 0; ii < 40; ii++ {
		test = DoStuff(test)
		fmt.Println("Iteration: ", ii+1, "Len: ", len(test))
	}

	fmt.Println("The Result length after 40 itterations is: ", len(test))

	for ii := 0; ii < 40; ii++ {
		test = DoStuff(test)
		fmt.Println("Iteration: ", ii+41, "Len: ", len(test))
	}

	fmt.Println("The Result length after 50 itterations is: ", len(test))
}

func DoStuff(numbers string) (result string) {
	counter := 0
	var num rune = rune(numbers[0])

	for _, number := range numbers {
		if number != num {
			result += fmt.Sprint(counter) + string(num)
			num = number
			counter = 1
		} else {
			counter += 1
		}
	}
	result += fmt.Sprint(counter) + string(num)
	return
}
