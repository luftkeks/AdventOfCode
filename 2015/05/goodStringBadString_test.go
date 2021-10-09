package main

import "testing"

func TestBehavingStrings(t *testing.T) {
	CreateRegexp()

	t.Run("ugknbfddgicrmopn", func(t *testing.T) {
		test := "ugknbfddgicrmopn"

		answer := IsStringGood(test)

		assertBool(t, answer, true)
	})
	t.Run("aaa", func(t *testing.T) {
		test := "aaa"

		answer := IsStringGood(test)

		assertBool(t, answer, true)
	})
	t.Run("jchzalrnumimnmhp", func(t *testing.T) {
		test := "jchzalrnumimnmhp"

		answer := IsStringGood(test)

		assertBool(t, answer, false)
	})
	t.Run("haegwjzuvuyypxyu", func(t *testing.T) {
		test := "haegwjzuvuyypxyu"

		answer := IsStringGood(test)

		assertBool(t, answer, false)
	})
	t.Run("dvszwmarrgswjxmb", func(t *testing.T) {
		test := "dvszwmarrgswjxmb"

		answer := IsStringGood(test)

		assertBool(t, answer, false)
	})
}

func assertBool(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
