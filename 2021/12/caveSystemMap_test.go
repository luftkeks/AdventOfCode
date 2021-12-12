package main

import (
	"testing"
)

func TestDoStuff(t *testing.T) {
	t.Run("First File", func(t *testing.T) {
		test := loadNodes("test1.txt")

		answer := Part1(test)

		assertInt(t, answer, 10)
	})
	t.Run("Second File", func(t *testing.T) {
		test := loadNodes("test2.txt")

		answer := Part1(test)

		assertInt(t, answer, 19)
	})
	t.Run("Third File", func(t *testing.T) {
		test := loadNodes("test3.txt")

		answer := Part1(test)

		assertInt(t, answer, 226)
	})
	t.Run("First File Part 2", func(t *testing.T) {
		test := loadNodes("test1.txt")

		answer := Part2(test)

		assertInt(t, answer, 36)
	})
	t.Run("Second File Part 2", func(t *testing.T) {
		test := loadNodes("test2.txt")

		answer := Part2(test)

		assertInt(t, answer, 103)
	})
	t.Run("Third File Part 2", func(t *testing.T) {
		test := loadNodes("test3.txt")

		answer := Part2(test)

		assertInt(t, answer, 3509)
	})
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
