package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	obsSum := Add(2, 3)
	expSum := 5

	if obsSum != expSum {
		t.Errorf("Observed sum = %d but expected sum = %d", obsSum, expSum)
	}
}

func ExampleAdd() {
	sum := Add(3, 6)
	fmt.Println(sum)
	// Output: 9
}
