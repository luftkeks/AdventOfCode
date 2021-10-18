package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Hardcode all the ingredients (as a struct) --> dont hardcode just read in the input
// make a cookie which can hold the ingredients
// make a function which calculates the score which takes a channel
// do a for for for for loop which gives all the possible combinations up to 100 --> can we generalize this?
// fire the function with the cookie channel
// check the output of the channel for score - save the highest scoring.

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

const (
	cookieCapacity int = 100
)

func main() {
	dat, _ := os.Open("input.txt")
	defer func(dat *os.File) { dat.Close() }(dat)
	scanner := bufio.NewScanner(dat)
	scannedStrings := []string{}
	for scanner.Scan() {
		scannedStrings = append(scannedStrings, scanner.Text())
	}

	ingridients := parseIngredients(scannedStrings)

	bestScore := 0
	for in1 := 0; in1 <= cookieCapacity; in1++ {
		for in2 := 0; in2 <= cookieCapacity-in1; in2++ {
			for in3 := 0; in3 <= cookieCapacity-in1-in2; in3++ {
				in4 := cookieCapacity - in1 - in2 - in3
				cookieScore := calculateCookie(ingridients, in1, in2, in3, in4, false)
				if cookieScore > bestScore {
					bestScore = cookieScore
				}
			}
		}
	}

	fmt.Printf("The best score is %v !\n", bestScore)
	bestScore = 0
	for in1 := 0; in1 <= cookieCapacity; in1++ {
		for in2 := 0; in2 <= cookieCapacity-in1; in2++ {
			for in3 := 0; in3 <= cookieCapacity-in1-in2; in3++ {
				in4 := cookieCapacity - in1 - in2 - in3
				cookieScore := calculateCookie(ingridients, in1, in2, in3, in4, true)
				if cookieScore > bestScore {
					bestScore = cookieScore
				}
			}
		}
	}
	fmt.Printf("The best score with 500 calories is %v !\n", bestScore)
}

func parseIngredients(lines []string) []Ingredient {
	ingridients := []Ingredient{}
	for _, line := range lines {
		words := strings.Split(line, " ")
		capacity, err := strconv.Atoi(words[2][:len(words[2])-1])
		if err != nil {
			log.Fatalln(err)
		}
		durability, err := strconv.Atoi(words[4][:len(words[4])-1])
		if err != nil {
			log.Fatalln(err)
		}
		flavor, err := strconv.Atoi(words[6][:len(words[6])-1])
		if err != nil {
			log.Fatalln(err)
		}
		texture, err := strconv.Atoi(words[8][:len(words[8])-1])
		if err != nil {
			log.Fatalln(err)
		}
		calories, err := strconv.Atoi(words[10][:len(words[10])])
		if err != nil {
			log.Fatalln(err)
		}
		ingridients = append(ingridients, Ingredient{name: words[0][:len(words[0])-1], capacity: capacity, durability: durability, flavor: flavor, texture: texture, calories: calories})
	}
	return ingridients
}

func calculateCookie(indrigients []Ingredient, in1, in2, in3, in4 int, caloriesLimit bool) int {
	totalCapacity := indrigients[0].capacity*in1 + indrigients[1].capacity*in2 + indrigients[2].capacity*in3 + indrigients[3].capacity*in4
	totalDurability := indrigients[0].durability*in1 + indrigients[1].durability*in2 + indrigients[2].durability*in3 + indrigients[3].durability*in4
	totalFlavor := indrigients[0].flavor*in1 + indrigients[1].flavor*in2 + indrigients[2].flavor*in3 + indrigients[3].flavor*in4
	totalTexture := indrigients[0].texture*in1 + indrigients[1].texture*in2 + indrigients[2].texture*in3 + indrigients[3].texture*in4
	if totalCapacity < 0 {
		totalCapacity = 0
	}
	if totalDurability < 0 {
		totalDurability = 0
	}
	if totalFlavor < 0 {
		totalFlavor = 0
	}
	if totalTexture < 0 {
		totalTexture = 0
	}
	totalCalories := indrigients[0].calories*in1 + indrigients[1].calories*in2 + indrigients[2].calories*in3 + indrigients[3].calories*in4
	if totalCalories != 500 && caloriesLimit {
		return 0
	}
	return totalCapacity * totalDurability * totalFlavor * totalTexture
}
