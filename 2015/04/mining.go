package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func main() {
	commandLineArgs := os.Args
	secret := commandLineArgs[1]

	low1, low2 := Mine(secret)

	fmt.Printf("The lowest Number with 5 trailing 0 in hash for the given secret is: %d", low1)
	fmt.Printf("The lowest Number with 6 trailing 0 in hash for the given secret is: %d", low2)
}

func Mine(secret string) (lowestNumber1, lowestNumber2 int) {
	for number := 0; lowestNumber2 == 0; number++ {
		hash := calculateHash(secret + strconv.Itoa(number))
		if hash[0:5] == "00000" && lowestNumber1 == 0 {
			lowestNumber1 = number
		}
		if hash[0:6] == "000000" {
			lowestNumber2 = number
			break
		}
	}
	return
}

func calculateHash(secret string) (hash string) {
	data := []byte(secret)
	hexHash := md5.Sum(data)
	return hex.EncodeToString(hexHash[:])
}
