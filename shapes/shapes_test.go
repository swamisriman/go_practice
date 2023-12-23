package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, name string, expected float64) {
		t.Helper()
		actual := shape.Area()

		if expected != actual {
			t.Errorf("shape: %q expected area: %g, but actual: %g", name, expected, actual)
		}
	}

	checkPerimeter := func(t testing.TB, shape Shape, name string, expected float64) {
		t.Helper()
		actual := shape.Perimeter()

		if expected != actual {
			t.Errorf("shape: %q expected perimeter: %g, but actual: %g", name, expected, actual)
		}
	}

	//Anonymous struct
	testData := []struct {
		name              string
		shape             Shape
		expectedArea      float64
		expectedPerimeter float64
	}{
		// Elements of the anonymous struct. As there is no name, there is nothing before '{'
		{name: "circle", shape: Circle{2.0}, expectedArea: math.Pi * 2 * 2, expectedPerimeter: 2 * math.Pi * 2},
		{"rectangle", Rectangle{2.0, 1.5}, 3.0, 7},
	}

	for _, datum := range testData {
		checkArea(t, datum.shape, datum.name, datum.expectedArea)
		checkPerimeter(t, datum.shape, datum.name, datum.expectedPerimeter)
	}
}
