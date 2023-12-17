package summer

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("any size", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		expectedSum := 10
		actualSum := Sum(arr)

		if expectedSum != actualSum {
			t.Errorf("Expected: %d but Actual: %d. Given: %v", expectedSum, actualSum, arr)
		}
	})
}

func TestSumAll(t *testing.T) {
	expectedSums := []int{15, 13}
	slicesInput := [][]int{{1, 2, 3, 4, 5}, {6, 7}}
	actualSums := SumAll(slicesInput[0], slicesInput[1])

	if !slices.Equal(expectedSums, actualSums) {
		t.Errorf("Expected: %v but Actual: %d. Given: %v", expectedSums, actualSums, slicesInput)
	}
}

func TestSumAllTails(t *testing.T) {
	assertSlicesEquality := func(t testing.TB, expected, actual []int) {
		t.Helper()
		if !slices.Equal(expected, actual) {
			t.Errorf("Expected: %v but Actual: %d", expected, actual)
		}
	}
	t.Run("Tails Sum, non-empty slice", func(t *testing.T) {
		expectedSums := []int{14, 7}
		slices := [][]int{{1, 2, 3, 4, 5}, {6, 7}}
		actualSums := SumAllTails(slices[0], slices[1])

		assertSlicesEquality(t, expectedSums, actualSums)
	})

	t.Run("Tails Sum, with empty slice", func(t *testing.T) {
		expectedSums := []int{14, 0}
		slices := [][]int{{1, 2, 3, 4, 5}, {}}
		actualSums := SumAllTails(slices[0], slices[1])

		assertSlicesEquality(t, expectedSums, actualSums)
	})
}
