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

type Literal struct {
	version int
	id      int
	number  int
}

type Operator struct {
	version int
	id      int
	content []Paket
}

type Paket interface {
	getID() int
	getVersion() int
	getVersionSum() int
	getLiteralSum() int
}

func main() {
	defer elapsed()()
	dat, err := os.Open("input.txt")
	if err != nil {
		panic("Hilfe File tut nicht")
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	scanner.Scan()
	inputString := scanner.Text()
	ints := ""
	for ii := 0; ii < len(inputString); ii++ {
		inInt, err := strconv.ParseInt(string(inputString[ii]), 16, 8)
		if err != nil {
			panic("Hilfe Zahl tut nicht")
		}
		ints = ints + fmt.Sprintf("%04b", inInt)
	}
	pack, lastBit := stringToPaket(ints)
	fmt.Println(lastBit)
	fmt.Println(pack.getLiteralSum())
	fmt.Printf("The Version Sum of all is: %v\n", pack.getVersionSum())
}

func stringToPaket(input string) (Paket, string) {
	version, _ := strconv.ParseInt(string(input[0:3]), 2, 8)
	id, _ := strconv.ParseInt(string(input[3:6]), 2, 8)
	var result Paket
	subString := string(input[7:])
	if id == 4 { //Literal
		content := ""
		for ii := 0; true; ii++ {
			startNextPaket := 11 + ii*5
			if len(input) > startNextPaket {
				subString = string(input[startNextPaket:])
			} else {
				subString = ""
			}
			if input[6+ii*5] == '1' {
				con := string(input[7+ii*5 : 11+ii*5])
				content += con
			} else {
				con := string(input[7+ii*5 : 11+ii*5])
				content += con
				break
			}
		}
		contentInt, _ := strconv.ParseInt(content, 2, 64)
		result = &Literal{version: int(version), id: int(id), number: int(contentInt)}
	} else { //Operator
		if input[6] == '1' {
			numberOfSubPackages, _ := strconv.ParseInt(string(input[7:18]), 2, 64)
			number := int(numberOfSubPackages)
			subPakets := []Paket{}
			subString = string(input[18:])
			var subPack Paket
			for ii := 0; ii < number; ii++ {
				subPack, subString = stringToPaket(subString)
				subPakets = append(subPakets, subPack)
			}
			result = &Operator{version: int(version), id: int(id), content: subPakets}
		} else {
			lengthOfSubPackages, _ := strconv.ParseInt(string(input[7:22]), 2, 64)
			length := int(lengthOfSubPackages)
			subPackages := []Paket{}
			subString2 := string(input[22 : 22+length])
			var subPack Paket
			for len(subString2) > 0 {
				subPack, subString2 = stringToPaket(subString2)
				subPackages = append(subPackages, subPack)
			}
			subString = string(input[22+length:])
			result = &Operator{version: int(version), id: int(id), content: subPackages}
		}
	}
	return result, subString
}

func (l *Literal) getVersion() int {
	return l.version
}

func (o *Operator) getVersion() int {
	return o.version
}

func (l *Literal) getID() int {
	return l.id
}

func (o *Operator) getID() int {
	return o.id
}

func (l *Literal) getVersionSum() int {
	return l.version
}

func (o *Operator) getVersionSum() int {
	result := o.version
	for _, pack := range o.content {
		result += pack.getVersionSum()
	}
	return result
}

func (l *Literal) getLiteralSum() int {
	return l.number
}

func (o *Operator) getLiteralSum() int {
	result := 0
	for _, pack := range o.content {
		result += pack.getLiteralSum()
	}
	return result
}

func (l *Literal) String() string {
	return fmt.Sprintf("%v %v", l.version, l.id)
}
func (o *Operator) String() string {
	return fmt.Sprintf("%v %v", o.version, o.id)
}
