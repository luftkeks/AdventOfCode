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

	setOfAll := map[Position]bool{}
	sensors[0].pos = Position{x: 0, y: 0, z: 0}
	for _, value := range sensors[0].beacons {
		setOfAll[value] = true
	}

	// vergleiche jeden relativvektor mit jedem - bei treffer count to 12 - if true use the rotation
	done := []int{0}
	for {
		for ii := 1; ii < len(sensors); ii++ {
			if !containsInt(done, ii) {
				hit := checkIfSensorIsAdjectedToAllMap(&setOfAll, sensors[ii])
				if hit {
					done = append(done, ii)
					fmt.Printf("Found position of sensor %v. In total we found %v of %v sensors.\n", ii, len(done), len(sensors))
				}
			}
		}
		if len(done) == len(sensors) {
			break
		}
	}
	fmt.Printf("There are %v unique beacons.\n", len(setOfAll))

	maxManhattanDistance := 0
	for ii := range sensors {
		for jj := 0; jj < ii; jj++ {
			manDistance := manhattanDistance(sensors[ii].pos, sensors[jj].pos)
			if maxManhattanDistance < manDistance {
				maxManhattanDistance = manDistance
			}
		}
	}
	fmt.Printf("The max Manhattan distance is: %v\n", maxManhattanDistance)
}

func checkIfSensorIsAdjectedToAllMap(maap *map[Position]bool, sensor *Sensor) bool {
	// get all list from map
	allList := getSliceFromMap(maap)
	// All list relativ to point schleife - save relativ point
	for index, fixPoint := range allList {
		listAllToCompare := getSliceRelativeToPoint(allList, index)
		// turn schleife x
		for xRot := 0; xRot < 4; xRot++ {
			// turn schleife y
			for yRot := 0; yRot < 4; yRot++ {
				// turn schleife z
				for zRot := 0; zRot < 4; zRot++ {
					// check list relative to point schleife - save sensor position
					for ii, matchingPoint := range sensor.beacons {
						listSensor := rotSlice(sensor.beacons, xRot, yRot, zRot)
						listSensorToCompare := getSliceRelativeToPoint(listSensor, ii)
						counter := 0
						// for point in checkListe if allList contains point counter++
						for _, item := range listSensorToCompare {
							if contains(listAllToCompare, item) {
								counter++
								// if counter >= 12
								if counter >= 12 {
									//sensor Calculation
									// get sensor rotation - get realtiv position of sensor
									(*sensor).pos = fixPoint.subtract(rotPos(matchingPoint, xRot, yRot, zRot))
									(*sensor).xRot = xRot
									(*sensor).yRot = yRot
									(*sensor).zRot = zRot
									// put everything together in all map.
									for _, elem := range listSensor {
										newBeaconPos := (*sensor).pos.add(elem)
										(*maap)[newBeaconPos] = true
									}
									return true
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func contains(positions []Position, other Position) bool {
	for _, value := range positions {
		if value == other {
			return true
		}
	}
	return false
}

func manhattanDistance(pos1, pos2 Position) int {
	return Abs(pos1.x-pos2.x) + Abs(pos1.y+pos2.y) + Abs(pos1.z+pos2.z)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func containsInt(positions []int, other int) bool {
	for _, value := range positions {
		if value == other {
			return true
		}
	}
	return false
}

func rotSlice(slice []Position, xRot, yRot, zRot int) []Position {
	result := make([]Position, len(slice))
	for index, pos := range slice {
		result[index] = rotPos(pos, xRot, yRot, zRot)
	}
	return result
}

func rotPos(pos Position, xRot, yRot, zRot int) Position {
	for xx := 0; xx < xRot; xx++ {
		pos = pos.turnX()
	}
	for yy := 0; yy < yRot; yy++ {
		pos = pos.turnY()
	}
	for zz := 0; zz < zRot; zz++ {
		pos = pos.turnZ()
	}
	return pos
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

func getSliceFromMap(maap *map[Position]bool) []Position {
	result := []Position{}
	for value := range *maap {
		result = append(result, value)
	}
	return result
}

func (p *Position) getRelativePosition(other Position) Position {
	return Position{x: other.x - p.x, y: other.y - p.y, z: other.z - p.z}
}

func (p *Position) add(other Position) Position {
	return Position{x: p.x + other.x, y: p.y + other.y, z: p.z + other.z}
}

func (p *Position) subtract(other Position) Position {
	return Position{x: p.x - other.x, y: p.y - other.y, z: p.z - other.z}
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
	var activeSensor *Sensor
	for _, line := range lines {
		if strings.Contains(line, "scanner") {
			activeSensor = &Sensor{name: line[4 : len(line)-4], beacons: []Position{}}
			sensors = append(sensors, activeSensor)
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
