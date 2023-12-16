package summer

import "testing"

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	expectedSum := 15
	actualSum := Sum(arr)

	if expectedSum != actualSum {
		t.Errorf("Expected: %d but Actual: %d. Given array %v", expectedSum, actualSum, arr)
	}

	t.Run("any size", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		expectedSum := 10
		actualSum := Sum(arr)

		if expectedSum != actualSum {
			t.Errorf("Expected: %d but Actual: %d. Given array %v", expectedSum, actualSum, arr)
		}
	})
}
