package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}

type Throw struct {
	initialVeloX, initialVeloY int
	velX, velY                 int
	posX, posY                 int
	highestY                   int
	hasHit                     bool
}

var targetStartX int
var targetFinishX int
var targetStartY int
var targetFinishY int

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
	inStuff := strings.Split(inputString, " ")
	xSplit := strings.Split(inStuff[2], "..")
	targetStartX, _ = strconv.Atoi(xSplit[0][2:])
	targetFinishX, _ = strconv.Atoi(xSplit[1][:len(xSplit[1])-1])
	ySplit := strings.Split(inStuff[3], "..")
	targetStartY, _ = strconv.Atoi(ySplit[0][2:])
	targetFinishY, _ = strconv.Atoi(ySplit[1])
	fmt.Println(targetStartX, targetFinishX, targetStartY, targetFinishY)

	throws := []*Throw{}
	wait := sync.WaitGroup{}
	for xx := 0; xx <= targetFinishX; xx++ {
		for yy := targetStartY; yy < 50000; yy++ {
			wait.Add(1)
			thr := Throw{velX: xx, velY: yy, initialVeloX: xx, initialVeloY: yy}
			throws = append(throws, &thr)
			go thr.throw(&wait)
		}
	}

	wait.Wait()
	highestY := 0
	counter := 0
	for _, threw := range throws {
		if highestY < threw.highestY {
			highestY = threw.highestY
		}
		if threw.hasHit {
			counter++
		}
	}

	fmt.Printf("The highest reached Y is: %v\n", highestY)
	fmt.Printf("The distinct velocities that reach target are: %v\n", counter)
}

func (t *Throw) throw(wait *sync.WaitGroup) {
	running := true
	for running {
		t.step()
		hit, err := t.checkIfTargetReached()
		if hit {
			t.hasHit = true
			break
		} else if err != nil {
			t.highestY = 0
			break
		}
	}
	wait.Done()
}

func (t *Throw) step() {
	t.posX += t.velX
	t.posY += t.velY
	if t.posY > t.highestY {
		t.highestY = t.posY
	}
	t.velY--
	if t.velX > 0 {
		t.velX--
	} else if t.velX < 0 {
		t.velX++
	}
}

func (t *Throw) checkIfTargetReached() (bool, error) {
	if targetFinishX < t.posX || t.posY < targetStartY {
		return false, Overshoot{}
	} else if targetStartX <= t.posX && t.posX <= targetFinishX && targetStartY <= t.posY && t.posY <= targetFinishY {
		return true, nil
	} else {
		return false, nil
	}
}

type Overshoot struct {
}

func (o Overshoot) Error() string {
	return "Das war zu weit."
}
