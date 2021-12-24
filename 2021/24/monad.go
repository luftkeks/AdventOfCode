package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	div := []int{}
	for line := 4; line < len(lines); line += 18 {
		ln := strings.Split(lines[line], " ")
		d, _ := strconv.Atoi(ln[2])
		div = append(div, d)
	}

	xs := []int{}
	for line := 5; line < len(lines); line += 18 {
		ln := strings.Split(lines[line], " ")
		x, _ := strconv.Atoi(ln[2])
		xs = append(xs, x)
	}

	ys := []int{}
	for line := 15; line < len(lines); line += 18 {
		ln := strings.Split(lines[line], " ")
		y, _ := strconv.Atoi(ln[2])
		ys = append(ys, y)
	}

	fmt.Printf("Divs: %v\n", printSlice(div))
	fmt.Printf("Xs  : %v\n", printSlice(xs))
	fmt.Printf("Ys  : %v\n", printSlice(ys))

	var input []int

	// start := time.Now()
	// for i1 := 9; i1 >= 1; i1-- {
	// 	for i2 := 9; i2 >= 1; i2-- {
	// 		for i3 := 9; i3 >= 1; i3-- {
	// 			for i4 := 9; i4 >= 1; i4-- {
	// 				for i5 := 9; i5 >= 1; i5-- {
	// 					for i6 := 9; i6 >= 1; i6-- {
	// 						for i7 := 9; i7 >= 1; i7-- {
	// 							for i8 := 9; i8 >= 1; i8-- {
	// 								for i9 := 9; i9 >= 1; i9-- {
	// 									for i10 := 9; i10 >= 1; i10-- {
	// 										for i11 := 9; i11 >= 1; i11-- {
	// 											for i12 := 9; i12 >= 1; i12-- {
	// 												for i13 := 9; i13 >= 1; i13-- {
	// 													for i14 := 9; i14 >= 1; i14-- {
	// 														input = []int{i1, i2, i4, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14}
	// 														z := calcZ(&xs, &ys, &div, &input)
	// 														if z == 0 {
	// 															fmt.Printf("Z Solution for %v is: %v.\n", input, z)
	// 															return
	// 														}
	// 													}
	// 												}
	// 											}
	// 										}
	// 									}
	// 								}
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}
	// 			fmt.Println(input)
	// 		}
	// 		fmt.Printf("Time since Start: %v\n", time.Since(start))
	// 	}
	// }

	input = []int{9, 3, 1, 8, 5, 1, 1, 1, 1, 2, 7, 9, 1, 1}
	fmt.Printf("inp : %v\n", printSlice(input))
	z := calcZ(&xs, &ys, &div, &input)
	fmt.Printf("\n")
	fmt.Printf("Z Solution for %v is: %v.\n", input, z)
}

func calcZ(xs, ys, div *[]int, input *[]int) (z int) {
	fmt.Printf("x   :   ")
	x, y, z := 0, 0, 0
	for index, val := range *input {
		x = z
		x = x % 26
		z /= (*div)[index]
		x += (*xs)[index]
		y = 1
		fmt.Printf("%v\t", x)
		if x != val {
			y += 25
			x = 1
		} else {
			x = 0
		}
		z = y * z
		y = val + (*ys)[index]
		y *= x
		z = z + y
	}
	return
}

func printSlice(slice []int) string {
	result := "[\t"
	for _, val := range slice {
		result += fmt.Sprintf("%v\t", val)
	}
	result += "]"
	return result
}
