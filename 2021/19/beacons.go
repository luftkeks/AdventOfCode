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

// This should turn into every direction
func (p *Position) getOrientation(ori int) Position {
	switch ori % 24 {
	case 0:
		return Position{x: -p.x, y: p.y, z: p.z}
	case 1:
		return Position{x: p.x, y: -p.y, z: p.z}
	case 2:
		return Position{x: p.x, y: p.y, z: -p.z}
	case 3:
		return Position{x: -p.x, y: -p.y, z: p.z}
	case 4:
		return Position{x: p.x, y: -p.y, z: -p.z}
	case 5:
		return Position{x: -p.x, y: p.y, z: -p.z}
	case 6:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 7:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 8:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 9:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 10:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 11:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 12:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 13:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 14:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 15:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 16:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 17:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 18:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 19:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 20:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 21:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 22:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	case 23:
		return Position{x: -p.x, y: -p.y, z: -p.z}
	default:
		panic("This shouldn't be done.")
	}
}
