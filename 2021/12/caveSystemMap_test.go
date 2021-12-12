package main

import (
	"testing"
)

func TestDoStuff(t *testing.T) {
	t.Run("First File", func(t *testing.T) {
		test := loadNodes("test1.txt")

		answer1, answer2 := Part(test)

		assertInt(t, answer1, 10)
		assertInt(t, answer2, 36)
	})
	t.Run("Second File", func(t *testing.T) {
		test := loadNodes("test2.txt")

		answer1, answer2 := Part(test)

		assertInt(t, answer1, 19)
		assertInt(t, answer2, 103)
	})
	t.Run("Third File", func(t *testing.T) {
		test := loadNodes("test3.txt")

		answer1, answer2 := Part(test)

		assertInt(t, answer1, 226)
		assertInt(t, answer2, 3509)
	})
}

func Benchmark_solution(t *testing.B) {
	test := loadNodes("input.txt")

	t.Run("Input", func(t *testing.B) {
		solution1, solution2 := Part(test)

		assertInt(t, solution1, 4378)
		assertInt(t, solution2, 133621)
	})
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
