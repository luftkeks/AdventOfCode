package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	From, To string
	Distance int
}

type PossibleRoute struct {
	Waypoints     []string
	TotalDistance int
}

var Routes []Route

func main() {
	commandLineArgs := os.Args
	fileToRead := commandLineArgs[1]
	dat, _ := os.Open(fileToRead)
	scanner := bufio.NewScanner(dat)

	Routes = []Route{}
	places := map[string]int{}
	for scanner.Scan() {
		route := parseLine(scanner.Text())
		Routes = append(Routes, route)
		places[route.From] = 1
		places[route.To] = 1
	}

	placeSlice := []string{}
	for place := range places {
		placeSlice = append(placeSlice, place)
	}

	//listPossibleRoutes := []PossibleRoute{}

	shortestRoute := PossibleRoute{TotalDistance: int(^uint(0) >> 1)}
	longestRoute := PossibleRoute{TotalDistance: 0}
	for _, permutation := range permutations(placeSlice) {
		route := PossibleRoute{Waypoints: permutation, TotalDistance: findTotalDistance(permutation)}
		//listPossibleRoutes = append(listPossibleRoutes, route)
		if route.TotalDistance < shortestRoute.TotalDistance {
			shortestRoute = route
		}
		if route.TotalDistance > longestRoute.TotalDistance {
			longestRoute = route
		}
	}

	fmt.Println("Shortest Route:" + shortestRoute.String())
	fmt.Println("Longest Route:" + longestRoute.String())
}

func parseLine(line string) (way Route) {
	strings := strings.Split(line, " ")
	distance, _ := strconv.Atoi(strings[4])
	return Route{From: strings[0], To: strings[2], Distance: distance}
}

func findTotalDistance(list []string) (distance int) {
	for ii := 0; ii < len(list)-1; ii++ {
		distance += findDistance(list[ii], list[ii+1])
	}
	return
}

func findDistance(from, to string) int {
	for _, route := range Routes {
		if (route.From == from && route.To == to) || (route.From == to && route.To == from) {
			return route.Distance
		}
	}
	panic("No Route found")
}

// Not my function - copied from: https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func (p *PossibleRoute) String() string {
	return "Waypoints: " + fmt.Sprintf("%v", p.Waypoints) + "Distance: " + fmt.Sprintf("%v", p.TotalDistance)
}
