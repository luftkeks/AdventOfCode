package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Box struct {
	Width, Height, Length int
}

func main() {
	commandLineArgs := os.Args
	fileToRead := commandLineArgs[1]

	dat, err := os.Open(fileToRead)
	check(err)

	scanner := bufio.NewScanner(dat)

	boxSlice := []Box{}
	for scanner.Scan() {
		boxSlice = append(boxSlice, createBoxFromString(scanner.Text()))
	}

	var totalPaper int
	var totalRibbon int

	for _, box := range boxSlice {
		totalPaper += box.calculatePaperNeeded()
		fmt.Println(box)
		fmt.Println(box.calculateRibbonNeeded())
		totalRibbon += box.calculateRibbonNeeded()
	}

	fmt.Printf("Amout of paper needed: ")
	fmt.Println(totalPaper)

	fmt.Printf("Amout of ribbon needed: ")
	fmt.Println(totalRibbon)
}

func (b Box) calculatePaperNeeded() (paperNedded int) {
	sides := []int{b.Width * b.Length, b.Width * b.Height, b.Height * b.Length}
	sort.Ints(sides)

	sides = append(sides, sides...)
	sides = append(sides, sides[0])
	for _, side := range sides {
		paperNedded += side
	}
	return
}

func (b Box) calculateRibbonNeeded() (ribbon int) {
	sides := []int{b.Height, b.Width, b.Length}
	sort.Ints(sides)
	return 2*sides[0] + 2*sides[1] + b.calculateVolume()
}

func (b Box) calculateVolume() (volume int) {
	return b.Width * b.Height * b.Length
}

func createBoxFromString(input string) Box {
	inputString := strings.Split(input, "x")
	length, _ := strconv.Atoi(inputString[0])
	width, _ := strconv.Atoi(inputString[1])
	height, _ := strconv.Atoi(inputString[2])
	return Box{Width: width, Height: height, Length: length}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
