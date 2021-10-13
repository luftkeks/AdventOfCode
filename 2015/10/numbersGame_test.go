package main

import (
	"fmt"
	"testing"
)

func TestDoStuff(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		test := "1"

		answer := DoStuff(test)

		assertString(t, answer, fmt.Sprint(11))
	})
	t.Run("11", func(t *testing.T) {
		test := "11"

		answer := DoStuff(test)

		assertString(t, answer, fmt.Sprint(21))
	})
	t.Run("21", func(t *testing.T) {
		test := "21"

		answer := DoStuff(test)

		assertString(t, answer, fmt.Sprint(1211))
	})
	t.Run("1211", func(t *testing.T) {
		test := "1211"

		answer := DoStuff(test)

		assertString(t, answer, fmt.Sprint(111221))
	})
	t.Run("111221", func(t *testing.T) {
		test := "111221"

		answer := DoStuff(test)

		assertString(t, answer, fmt.Sprint(312211))
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
