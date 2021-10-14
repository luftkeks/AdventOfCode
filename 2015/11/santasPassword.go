package main

import (
	"bytes"
	"fmt"
	"strings"
)

var AbcList []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var AbcdList []byte = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

type Password struct {
	str string
}

func (p Password) testPassword() bool {
	return p.TestFirstRequirement() && p.TestSecondRequirement() && p.TestThirdRequirement()
}

func (p Password) TestFirstRequirement() bool {
	for ii := 0; ii < len(AbcList)-2; ii++ {
		if strings.Contains(p.str, AbcList[ii]+AbcList[ii+1]+AbcList[ii+2]) {
			return true
		}
	}
	return false
}

func (p Password) TestSecondRequirement() bool {
	if strings.Contains(p.str, "l") || strings.Contains(p.str, "i") || strings.Contains(p.str, "o") {
		return false
	}
	return true
}

func (p Password) TestThirdRequirement() bool {

	numberOfPairs := 0
	for _, char := range AbcList {
		if strings.Contains(p.str, char+char) {
			numberOfPairs++
		}
	}
	return numberOfPairs >= 2
}

func (p *Password) countOneUp(stage int) {
	password := p.str
	charNow := password[len(p.str)-stage-1]
	char := bytes.IndexByte(AbcdList, charNow)
	if char+1 < len(AbcdList) {
		p.str = changeCharInPassword(p.str, AbcList[char+1], len(p.str)-stage)
	} else {
		p.str = changeCharInPassword(p.str, string(AbcList[0]), len(p.str)-stage)
		p.countOneUp(stage + 1)
	}
}

func changeCharInPassword(password string, char string, position int) string {
	var passNew string
	if position > 0 && position < len(password) {
		passNew = password[:position-1] + char + password[position:]
	} else if position == 0 {
		passNew = char + password[position+1:]
	} else if position == len(password) {
		passNew = password[:position-1] + char
	}
	return passNew
}

func main() {
	pass := Password{str: "hxbxwxba"}
	pass = getNextPassword(pass)
	fmt.Println("Santas next Password is: ", pass)
	pass = getNextPassword(pass)
	fmt.Println("And Santas Password afterwards is: ", pass)

}

func getNextPassword(pass Password) Password {
	for {
		pass.countOneUp(0)
		if pass.testPassword() {
			return pass
		}
	}
}
