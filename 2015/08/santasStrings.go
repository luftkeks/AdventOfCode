package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	commandLineArgs := os.Args
	fileToRead := commandLineArgs[1]
	input, err := ioutil.ReadFile(fileToRead)
	if err != nil {
		panic(err)
	}

	scannedStrings := strings.Split(string(input), "\n")

	numberOfLines := len(scannedStrings)
	lenStrings := 0
	lenUnquote := 0
	lenQuote := 0

	for ii := 0; ii < numberOfLines; ii++ {
		lenStrings += len(scannedStrings[ii])
		str, _ := strconv.Unquote(scannedStrings[ii])
		lenUnquote += len(str)
		lenQuote += len(strconv.Quote(scannedStrings[ii]))
		fmt.Printf("%s \t %s\n", scannedStrings[ii], strconv.Quote(scannedStrings[ii]))
	}

	fmt.Println("LenStrings", lenStrings)
	fmt.Println("LenUnquote", lenUnquote)
	fmt.Println("LenStrings - LenUnquote", lenStrings-lenUnquote)
	fmt.Println("lenUnquote - LenStrings", lenQuote-lenStrings)

}
