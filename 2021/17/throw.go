package main

import (
	"bufio"
	"fmt"
	"os"
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

var targetStartX int = 128
var targetFinishX int = 160
var targetStartY int = -142
var targetFinishY int = -88

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
	fmt.Println(inputString)

	throws := []*Throw{}
	wait := sync.WaitGroup{}
	for xx := 10; xx < 128; xx++ {
		for yy := 20; yy < 50000; yy++ {
			wait.Add(1)
			thr := Throw{velX: xx, velY: yy}
			throws = append(throws, &thr)
			go thr.throw(&wait)
		}
	}

	wait.Wait()
	highestY := 0
	for _, threw := range throws {
		if highestY < threw.highestY {
			highestY = threw.highestY
		}
	}

	fmt.Printf("The highest reached Y is: %v\n", highestY)
}

func (t *Throw) throw(wait *sync.WaitGroup) {
	running := true
	for running {
		t.step()
		hit, err := t.checkIfTargetReached()
		if hit {
			t.hasHit = true
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
