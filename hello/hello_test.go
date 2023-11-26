package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	var obs, exp string

	obs = Hello("Sriman")
	exp = "Hello Sriman"
	assertStringsEquality(t, exp, obs)

	//sub-test demonstration
	t.Run("Use 'World' when empty string is passed", func(t *testing.T) {
		obs = Hello("")
		exp = "Hello World"
		assertStringsEquality(t, exp, obs)
	})
}

func assertStringsEquality(t testing.TB, exp, obs string) {
	t.Helper()

	if exp == obs {
		fmt.Println("All good")
	} else {
		t.Errorf("Obtained: %q but Expected: %q", obs, exp)
	}
}
