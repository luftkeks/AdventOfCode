package main

import (
	"testing"
)

func TestSantasPassword(t *testing.T) {
	t.Run("hijklmmn", func(t *testing.T) {
		test := Password{"hijklmmn"}

		answer := test.TestFirstRequirement()
		assertBool(t, answer, true)
		answer = test.TestSecondRequirement()
		assertBool(t, answer, false)

	})
	t.Run("abbceffg", func(t *testing.T) {
		test := Password{"abbceffg"}

		answer := test.TestFirstRequirement()
		assertBool(t, answer, false)
		answer = test.TestThirdRequirement()
		assertBool(t, answer, true)

	})
	t.Run("abbcegjk", func(t *testing.T) {
		test := Password{"abbcegjk"}

		answer := test.TestThirdRequirement()
		assertBool(t, answer, false)
	})

}

func assertBool(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
