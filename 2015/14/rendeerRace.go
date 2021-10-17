package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Rendeer struct {
	name     string
	postion  int
	speed    int
	moveTime int
	waitTime int
	nowDoing string
	timeToDo int
	points   int
}

const (
	runtime int = 2503
)

func main() {
	dat, _ := os.Open("input.txt")
	defer func(dat *os.File) { dat.Close() }(dat)
	scanner := bufio.NewScanner(dat)
	scannedStrings := []string{}
	for scanner.Scan() {
		scannedStrings = append(scannedStrings, scanner.Text())
	}

	var rendeers []*Rendeer
	var rendeers2 []*Rendeer
	for _, line := range scannedStrings {
		ren := createRendeer(line)
		ren2 := createRendeer(line)
		rendeers = append(rendeers, &ren)
		rendeers2 = append(rendeers2, &ren2)
	}

	channel := make(chan bool)
	for _, rendeer := range rendeers {
		go rendeer.runTForTime(runtime, channel)
	}

	for ii := 0; ii < len(rendeers); ii++ {
		<-channel
	}
	sort.Slice(rendeers, func(i, j int) bool { return rendeers[i].postion > rendeers[j].postion })

	fmt.Printf("The leader in terms of position is %v, with position %v \n", rendeers[0].name, rendeers[0].postion)

	for ii := 0; ii < runtime; ii++ {
		for _, rendeer2 := range rendeers2 {
			rendeer2.run()
		}
		sort.Slice(rendeers2, func(i, j int) bool { return rendeers2[i].postion > rendeers2[j].postion })
		for _, ren2 := range rendeers2 {
			ren2.getPointForPosition(rendeers2[0].postion)
		}
	}

	sort.Slice(rendeers2, func(i, j int) bool { return rendeers2[i].points > rendeers2[j].points })
	fmt.Printf("The leader in terms of points is %v, with position %v and points %v \n", rendeers2[0].name, rendeers2[0].postion, rendeers2[0].points)
}

func createRendeer(line string) Rendeer {
	words := strings.Split(line, " ")
	speed, err := strconv.Atoi(words[3])
	if err != nil {
		log.Fatalln("Speed was not a number")
	}
	runTime, err := strconv.Atoi(words[6])
	if err != nil {
		log.Fatalln("Run time was not a number")
	}
	waitTime, err := strconv.Atoi(words[13])
	if err != nil {
		log.Fatalln("Wait time was not a number")
	}
	return Rendeer{name: words[0], postion: 0, speed: speed, moveTime: runTime, waitTime: waitTime, nowDoing: "run", timeToDo: runTime, points: 0}
}

func (r *Rendeer) getPointForPosition(pos int) {
	if r.postion == pos {
		r.points++
	}
}

func (r *Rendeer) runTForTime(time int, ready chan bool) {
	for ii := 0; ii < time; ii++ {
		r.run()
	}
	ready <- true
}

func (r *Rendeer) run() {
	if r.timeToDo == 0 {
		r.switchMovingType()
		r.run()
	} else {
		r.timeToDo--
		if r.nowDoing == "run" {
			r.postion += r.speed
		} else if r.nowDoing == "wait" {
			// do nothing
		} else {
			log.Fatalln("now moving is not waiting and not runnig")
		}
	}
}

func (r *Rendeer) switchMovingType() {
	if r.nowDoing == "wait" {
		r.nowDoing = "run"
		r.timeToDo = r.moveTime
	} else if r.nowDoing == "run" {
		r.nowDoing = "wait"
		r.timeToDo = r.waitTime
	} else {
		log.Fatalln("This should NOT happen")
	}
}
