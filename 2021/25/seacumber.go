package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func elapsed() func() {
	start := time.Now()
	return func() { fmt.Printf("Day took %v\n", time.Since(start)) }
}

func main() {
	defer elapsed()()
	dat, err := os.Open("input.txt")
	if err != nil {
		panic("Hilfe File tut nicht")
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	maap := [][]rune{}
	for scanner.Scan() {
		inString := scanner.Text()
		maap = append(maap, []rune(inString))
	}

	moved := true
	step := 0
	for i := 1; i < 5000 && moved; i++ {
		maap, moved = calculateStep(maap)
		step = i
	}

	fmt.Printf("Seacumber stoped moving after step: %v\n", step)

}

func calculateStep(mapp [][]rune) ([][]rune, bool) {
	mtempE, mtempS := true, true
	mapp, mtempE = calculateStepHordeEast(mapp)
	mapp, mtempS = calculateStepHordeSouth(mapp)
	return mapp, mtempE || mtempS
}

func calculateStepHordeSouth(mapp [][]rune) ([][]rune, bool) {
	moved := false
	result := createNewMap(mapp, '>')
	for yy := 0; yy < len(result); yy++ {
		for xx := 0; xx < len(result[yy]); xx++ {
			if mapp[yy][xx] == 'v' && mapp[(yy+1)%len(result)][xx] == '.' {
				result[(yy+1)%len(result)][xx] = 'v'
				moved = true
			} else if result[yy][xx] != 'v' {
				result[yy][xx] = mapp[yy][xx]
			}
		}
	}
	return result, moved
}

func calculateStepHordeEast(mapp [][]rune) ([][]rune, bool) {
	moved := false
	result := createNewMap(mapp, 'v')
	for yy := 0; yy < len(result); yy++ {
		for xx := 0; xx < len(result[yy]); xx++ {
			if mapp[yy][xx] == '>' && mapp[yy][(xx+1)%len(mapp[yy])] == '.' {
				result[yy][(xx+1)%len(mapp[yy])] = '>'
				moved = true
			} else if result[yy][xx] != '>' {
				result[yy][xx] = mapp[yy][xx]
			}
		}
	}
	return result, moved
}

func createNewMap(mapp [][]rune, keep ...rune) [][]rune {
	result := make([][]rune, len(mapp))
	for yy := 0; yy < len(result); yy++ {
		result[yy] = make([]rune, len(mapp[yy]))
		for xx := 0; xx < len(result[yy]); xx++ {
			if containsRune(keep, mapp[yy][xx]) {
				result[yy][xx] = mapp[yy][xx]
			} else {
				result[yy][xx] = '.'
			}
		}
	}
	return result
}

func printMap(mapp [][]rune) {
	fmt.Printf("\n")
	for ii := 0; ii < len(mapp); ii++ {
		fmt.Printf("%v\n", string(mapp[ii]))
	}
	fmt.Printf("\n")
}

func containsRune(slice []rune, other rune) bool {
	for _, value := range slice {
		if value == other {
			return true
		}
	}
	return false
}
