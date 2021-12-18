package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	defer elapsed()()

	filename := "input.txt"
	pair := parseInputOne(filename)

	fmt.Println(pair)
	fmt.Printf("The magnitude of the result is: %v\n", pair.calcMagnitude())

	fmt.Printf("The max Magnitude of two Lines is %v\n", secondPart(filename))
}

func secondPart(input string) int {
	dat, err := os.Open(input)
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

	maxMagnitude := 0
	for ii := 0; ii < len(lines); ii++ {
		for jj := 0; jj < len(lines); jj++ {
			if ii != jj {
				p1, _ := createDoppel([]rune(lines[ii]), 0)
				p2, _ := createDoppel([]rune(lines[jj]), 0)
				mag := reductDoppel(&Doppel{pair1: p1, pair2: p2}).calcMagnitude()
				if mag > maxMagnitude {
					maxMagnitude = mag
				}
			}
		}
	}
	return maxMagnitude
}

func parseInputOne(input string) Pair {
	dat, err := os.Open(input)
	if err != nil {
		panic("Hilfe File tut nicht")
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)

	scanner.Scan()
	inString := scanner.Text()
	pair, _ := createDoppel([]rune(inString), 0)
	pair = reductDoppel(pair)
	for scanner.Scan() {
		inString := scanner.Text()
		dop, _ := createDoppel([]rune(inString), 0)
		pairTemp := Doppel{pair1: pair, pair2: dop}
		pair = reductDoppel(&pairTemp)
	}
	return pair
}

func reductDoppel(pair Pair) Pair {
	needsReducted := true
	for needsReducted {
		explode, split := pair.hasError(0)
		if explode {
			pair.(*Doppel).reduct(0)
		} else if split {
			pair.split()
		} else {
			needsReducted = explode || split
		}
	}
	return pair
}

func createDoppel(in []rune, point int) (Pair, int) {
	if in[point] == '[' {
		dop1, dot1 := createDoppel(in, point+1)
		dop2, dot2 := createDoppel(in, dot1)
		return &Doppel{pair1: dop1, pair2: dop2}, dot2
	} else if in[point] >= '0' && in[point] <= '9' {
		return &Number{number: int(in[point] - '0')}, point + 1
	} else {
		return createDoppel(in, point+1)
	}
}

type Pair interface {
	addLeft(int)
	addRight(int)
	getNum() (int, bool)
	split() (Pair, bool)
	hasError(depth int) (bool, bool)
	calcMagnitude() int
}

type Number struct {
	number int
}

type Doppel struct {
	pair1, pair2 Pair
}

func (n *Number) calcMagnitude() int {
	return n.number
}

func (d *Doppel) calcMagnitude() int {
	return 3*d.pair1.calcMagnitude() + 2*d.pair2.calcMagnitude()
}

func (n *Number) hasError(depth int) (bool, bool) {
	return depth > 4, n.number > 9
}

func (d *Doppel) hasError(depth int) (bool, bool) {
	leftDeep, leftNum := d.pair1.hasError(depth + 1)
	rightDeep, rightNum := d.pair2.hasError(depth + 1)
	return leftDeep || rightDeep, leftNum || rightNum
}

func (n *Number) split() (result Pair, didSomething bool) {
	if n.number > 9 {
		result = &Doppel{pair1: &Number{number: n.number / 2}, pair2: &Number{number: n.number/2 + n.number%2}}
	} else {
		result = n
	}
	return result, n.number > 9
}

func (d *Doppel) split() (result Pair, didSome bool) {
	d.pair1, didSome = d.pair1.split()
	if !didSome {
		d.pair2, didSome = d.pair2.split()
	}
	return d, didSome
}

func (d *Doppel) explode() (int, int) {
	number1, _ := d.pair1.getNum()
	number2, _ := d.pair2.getNum()

	return number1, number2
}

func (d *Doppel) reduct(depth int) (number1, number2 int, didSomething bool) {
	_, is1Num := d.pair1.getNum()
	_, is2Num := d.pair2.getNum()
	if depth == 3 {
		if !is1Num {
			num1, num2 := d.pair1.(*Doppel).explode()
			d.pair1 = &Number{number: 0}
			d.pair2.addLeft(num2)
			number1 = num1
			didSomething = true
		}
		if !is2Num && !didSomething {
			num1, num2 := d.pair2.(*Doppel).explode()
			d.pair2 = &Number{number: 0}
			d.pair1.addRight(num1)
			number2 = num2
			didSomething = true
		}
	} else {
		if !is1Num {
			num1, num2, did := d.pair1.(*Doppel).reduct(depth + 1)
			d.pair2.addLeft(num2)
			number1 = num1
			didSomething = did
		}
		if !is2Num && !didSomething {
			num1, num2, did := d.pair2.(*Doppel).reduct(depth + 1)
			d.pair1.addRight(num1)
			number2 = num2
			didSomething = did
		}
	}
	return
}

func (d *Doppel) addLeft(num int) {
	d.pair1.addLeft(num)
}

func (d *Doppel) addRight(num int) {
	d.pair2.addRight(num)
}

func (n *Number) addLeft(num int) {
	n.number += num
}

func (n *Number) addRight(num int) {
	n.number += num
}

func (n *Number) getNum() (int, bool) {
	return n.number, true
}

func (d *Doppel) getNum() (int, bool) {
	return 0, false
}

func (d *Doppel) String() string {
	return fmt.Sprintf("[%v,%v]", d.pair1, d.pair2)
}
func (n *Number) String() string {
	return strconv.Itoa(n.number)
}

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}
