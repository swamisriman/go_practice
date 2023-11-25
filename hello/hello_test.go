package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	// exp := "Hello World"
	// obs := Hello()

	// if exp == obs {
	// 	fmt.Println("All good")
	// } else {
	// 	t.Errorf("Obtained: %q but Expected: %q", obs, exp)
	// }

	obs := Hello("Sriman")
	exp := "Hello Sriman"

	if exp == obs {
		fmt.Println("All good")
	} else {
		t.Errorf("Obtained: %q but Expected: %q", obs, exp)
	}
}
