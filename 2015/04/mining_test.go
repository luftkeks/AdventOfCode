package main

import "testing"

func TestMining(t *testing.T) {

	t.Run("abcdef", func(t *testing.T) {
		secret := "abcdef"

		lowestInt, _ := Mine(secret)

		assertInts(t, lowestInt, 609043)
	})

	t.Run("pqrstuv", func(t *testing.T) {
		secret := "pqrstuv"

		lowestInt, _ := Mine(secret)

		assertInts(t, lowestInt, 1048970)
	})
}

func assertInts(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
