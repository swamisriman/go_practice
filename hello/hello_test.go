package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	var obs, exp string

	t.Run("Language specific greeting - Telugu", func(t *testing.T) {
		obs = Hello("Telugu", "శ్రీమన్")
		exp = "నమస్తే శ్రీమన్"
		assertStringsEquality(t, exp, obs)
	})

	t.Run("Language specific greeting - Spanish", func(t *testing.T) {
		obs = Hello("Spanish", "Pablo")
		exp = "Hola Pablo"
		assertStringsEquality(t, exp, obs)
	})

	t.Run("Use 'World' when empty string is passed and use English when unknown language is passed", func(t *testing.T) {
		obs = Hello("", "")
		exp = "Hello World"
		assertStringsEquality(t, exp, obs)
	})

	t.Run("Use English when unknown language is passed", func(t *testing.T) {
		obs = Hello("asfd", "Sriman")
		exp = "Hello Sriman"
		assertStringsEquality(t, exp, obs)
	})

	t.Run("Language specific greeting - Telugu, when empty name is passed", func(t *testing.T) {
		obs = Hello("Telugu", "")
		exp = "నమస్తే World"
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
