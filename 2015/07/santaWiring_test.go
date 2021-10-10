package main

import (
	"testing"
)

func TestWiring(t *testing.T) {

	t.Run("Test", func(t *testing.T) {

		wiring := Wiring{}

		strings := []string{"123 -> x",
			"456 -> y",
			"x AND y -> d",
			"x OR y -> e",
			"x LSHIFT 2 -> f",
			"y RSHIFT 2 -> g",
			"NOT x -> h",
			"NOT y -> i"}
		WireMap(strings, wiring)

		for key, _ := range wiring {
			assertStuff(t, key, wiring.Get(key).Get(wiring))
		}
	})
}

func assertStuff(t testing.TB, key string, got uint16) {
	want := map[string]uint16{"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456}

	if got != want[key] {
		t.Errorf("key %v got %v want %v", key, got, want[key])
	}
}
