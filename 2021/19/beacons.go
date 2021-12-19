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
	name             string
	pos              Position
	beacons          []Position
	xRot, yRot, zRot int
}

func main() {
	defer elapsed()()
	sensors := readInSensors("test.txt")
	fmt.Println(sensors)

	setOfAll := map[Position]bool{}
	sensors[0].pos = Position{x: 0, y: 0, z: 0}
	for _, value := range sensors[0].beacons {
		setOfAll[value.getRelativePosition(sensors[0].beacons[0])] = true
	}
	sensors[0].pos = sensors[0].pos.getRelativePosition(sensors[0].beacons[0])

	// vergleiche jeden relativvektor mit jedem - bei treffer count to 12 - if true use the rotation

	// All list relativ to point schleife
	// turn schleife x
	// turn schleife y
	// turn schleife z
	// check list relative to point schleife
	// for point in checkListe if allList contains point counter++
	// if counter >= 12
	// get sensor rotation - get realtiv position of new sensor
	// put everything together in all list.
}

func contains(positions []Position, other Position) bool {
	for _, value := range positions {
		if value == other {
			return true
		}
	}
	return false
}

func getSliceRelativeToPoint(points []Position, num int) []Position {
	if num > len(points) {
		panic("This is more then i can take")
	}
	newPoints := []Position{}
	for _, value := range points {
		newPoints = append(newPoints, value.getRelativePosition(points[num]))
	}
	return newPoints
}

func (p *Position) getRelativePosition(other Position) Position {
	return Position{x: other.x - p.x, y: other.y - p.y, z: other.z - p.z}
}

func readInSensors(input string) []*Sensor {
	dat, err := os.Open(input)
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
	return sensors
}

// This should turn into every direction
func (p *Position) turnX() Position {
	return Position{x: p.x, y: p.z, z: -p.y}
}

func (p *Position) turnY() Position {
	return Position{x: p.z, y: p.y, z: -p.x}
}

func (p *Position) turnZ() Position {
	return Position{x: p.y, y: -p.x, z: p.z}
}
