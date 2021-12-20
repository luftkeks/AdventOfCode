package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}

func main() {
	defer elapsed()()
	dat, err := os.Open("input.txt")
	if err != nil {
		panic("Hilfe File tut nicht")
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	lines := []string{}
	for scanner.Scan() {
		inString := scanner.Text()
		lines = append(lines, inString)
	}

	lookup := []rune(lines[0])

	// create picture as char 2D Slice and guarante a border of '.'
	picture := make([][]rune, len(lines)-2)
	picture[0] = []rune{}
	for ii := 0; ii < len(lines)-2; ii++ {
		picture[ii] = []rune(lines[ii+2])
	}

	printPicture(picture)
	inf := '.'
	for jj := 0; jj < 2; jj++ {
		picture, inf = createPictureFromOld(picture, lookup, inf)
		printPicture(picture)
	}

	counter := 0
	for yy := 0; yy < len(picture); yy++ {
		for xx := 0; xx < len(picture[yy]); xx++ {
			if picture[yy][xx] == '#' {
				counter++
			}
		}
	}

	fmt.Printf("The number of lit things after two turns is: %v\n", counter)

	for jj := 0; jj < 48; jj++ {
		picture, inf = createPictureFromOld(picture, lookup, inf)
	}

	counter = 0
	for yy := 0; yy < len(picture); yy++ {
		for xx := 0; xx < len(picture[yy]); xx++ {
			if picture[yy][xx] == '#' {
				counter++
			}
		}
	}

	fmt.Printf("The number of lit things after 50 turns is: %v\n", counter)
}

// make border 2 wide
func createPictureFromOld(old [][]rune, lookup []rune, infinity rune) (new [][]rune, inf rune) {
	new = make([][]rune, len(old)+2)
	for ii := 0; ii < len(new); ii++ {
		new[ii] = make([]rune, len(old[0])+2)
	}

	for yy := 0; yy < len(new); yy++ {
		for xx := 0; xx < len(new[0]); xx++ {
			niner := make([][]rune, 3)
			for ii := -1; ii <= 1; ii++ {
				niner[ii+1] = make([]rune, 3)
				for jj := -1; jj <= 1; jj++ {
					if yy+ii-1 < 0 || yy+ii-1 >= len(old) || xx+jj-1 < 0 || xx+jj-1 >= len(old[0]) {
						niner[ii+1][jj+1] = infinity

					} else {
						niner[ii+1][jj+1] = old[yy+ii-1][xx+jj-1]
					}
				}
			}
			numberOfNiner := genIntFromRuneSquare(niner)
			new[yy][xx] = lookup[numberOfNiner]
		}
	}

	niner := make([][]rune, 3)
	for ii := 0; ii < 3; ii++ {
		niner[ii] = make([]rune, 3)
		for jj := 0; jj < 3; jj++ {
			niner[ii][jj] = infinity
		}
	}

	inf = lookup[genIntFromRuneSquare(niner)]
	return
}

func genIntFromRuneSquare(square [][]rune) int {
	str := []rune{}
	for yy := 0; yy < len(square); yy++ {
		for xx := 0; xx < len(square[0]); xx++ {
			if square[yy][xx] == '#' {
				str = append(str, '1')
			} else {
				str = append(str, '0')
			}
		}
	}
	return getIntFromBitString(string(str))
}

func getIntFromBitString(str string) int {
	number, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic("Irgendwas stimmt mit number nicht.")
	}
	return int(number)
}

func printPicture(pic [][]rune) {
	fmt.Printf("\n")
	for yy := 0; yy < len(pic); yy++ {
		fmt.Printf("%v\n", string(pic[yy]))
	}
	fmt.Printf("\n")
}
