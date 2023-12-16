package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	expected := "aaaa"
	actual := Repeat("a", 4)

	if expected != actual {
		t.Errorf("Expected: %q, but Actual: %q", expected, actual)
	}

}

func TestRepeatInBuilt(t *testing.T) {
	expected := strings.Repeat("a", 4)
	actual := Repeat("a", 4)

	if expected != actual {
		t.Errorf("Expected: %q, but Actual: %q", expected, actual)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	symbol := "ab"
	repetitionCount := 5
	repeatedSymbol := Repeat(symbol, repetitionCount)

	fmt.Println(repeatedSymbol)
	// Output: ababababab
}
