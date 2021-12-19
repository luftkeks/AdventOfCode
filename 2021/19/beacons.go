package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}

type Position struct {
	x, y, z int
}

type Sensor struct {
	name        string
	pos         Position
	orientation int
	beacons     []Position
}

func main() {
	defer elapsed()()

	dat, err := os.Open("test.txt")
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

	sensors := []*Sensor{}
	var activeSensor Sensor
	for _, line := range lines {
		if strings.Contains(line, "scanner") {
			activeSensor = Sensor{name: line[4 : len(line)-4], beacons: []Position{}}
			sensors = append(sensors, &activeSensor)
			continue
		} else if len(line) != 0 {
			pos := strings.Split(line, ",")
			x, _ := strconv.Atoi(pos[0])
			y, _ := strconv.Atoi(pos[1])
			z, _ := strconv.Atoi(pos[2])
			activeSensor.beacons = append(activeSensor.beacons, Position{x: x, y: y, z: z})
		}
	}
	fmt.Println(sensors)
}
