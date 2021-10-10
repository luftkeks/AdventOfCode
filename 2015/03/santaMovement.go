package main

import (
	"fmt"
	"os"
)

type Direction int64

const (
	Up    Direction = 0
	Down  Direction = 1
	Right Direction = 2
	Left  Direction = 3
)

type Point struct {
	x, y int
}

type SantasMap map[Point]bool

func main() {
	commandLineArgs := os.Args
	fileToRead := commandLineArgs[1]

	dat, err := os.ReadFile(fileToRead)
	check(err)

	wayToGo := string(dat)

	santaPostion := Point{x: 0, y: 0}

	card := SantasMap{}
	card.addPoint(santaPostion)

	for _, char := range wayToGo {
		switch char {
		case '<':
			move(Left, &santaPostion, &card)
		case '>':
			move(Right, &santaPostion, &card)
		case 'v':
			move(Down, &santaPostion, &card)
		case '^':
			move(Up, &santaPostion, &card)
		}
	}

	fmt.Printf("Number of visited Houses: ")
	fmt.Println(card.countPoints())

	santaPostion2 := Point{x: 0, y: 0}
	roboSantaPostion2 := Point{x: 0, y: 0}

	card2 := SantasMap{}
	card2.addPoint(santaPostion2)

	for number, char := range wayToGo {

		var dummyPosition *Point

		if number%2 == 0 {
			dummyPosition = &santaPostion2
		} else {
			dummyPosition = &roboSantaPostion2
		}

		switch char {
		case '<':
			move(Left, dummyPosition, &card2)
		case '>':
			move(Right, dummyPosition, &card2)
		case 'v':
			move(Down, dummyPosition, &card2)
		case '^':
			move(Up, dummyPosition, &card2)
		}
	}

	fmt.Printf("Number of visited Houses with Robo Santa: ")
	fmt.Println(card2.countPoints())
}

func move(direction Direction, santaPosition *Point, card *SantasMap) {
	switch direction {
	case Up:
		santaPosition.y += 1
	case Down:
		santaPosition.y -= 1
	case Right:
		santaPosition.x += 1
	case Left:
		santaPosition.x -= 1
	}

	card.addPoint(*santaPosition)
}

func (s SantasMap) addPoint(p Point) {
	s[p] = true
}

func (s SantasMap) countPoints() int {
	return len(s)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
