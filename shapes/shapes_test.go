package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Getting perimeter of rectangle when sides are given", func(t *testing.T) {
		rect := Rectangle{2.0, 5.0}
		actual := Perimeter(rect)
		expected := 14.0

		if actual != expected {
			t.Errorf("expected: %.2f, but actual: %2.f", expected, actual)
		}
	})

	t.Run("Getting perimeter of rectangle when sides are given using Rectangle struct's method", func(t *testing.T) {
		rect := Rectangle{2.0, 6.0}
		actual := rect.Perimmeter()
		expected := 16.0

		if actual != expected {
			t.Errorf("expected: %.2f, but actual: %2.f", expected, actual)
		}
	})

	t.Run("Getting perimeter of circle when radius is given using Circle struct's method", func(t *testing.T) {
		circle := Circle{2.0}
		actual := circle.Perimeter()
		expected := 2 * math.Pi * 2

		if actual != expected {
			t.Errorf("expected: %g, but actual: %g", expected, actual)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("Getting area of rectangle when sides are given", func(t *testing.T) {
		rect := Rectangle{2.0, 5.0}
		actual := Area(rect)
		expected := 10.0

		if actual != expected {
			t.Errorf("expected: %.2f, but actual: %2.f", expected, actual)
		}
	})

	t.Run("Getting area of rectangle when sides are given using Rectangle struct's method", func(t *testing.T) {
		rect := Rectangle{2.0, 6.0}
		actual := rect.Area()
		expected := 12.0

		if actual != expected {
			t.Errorf("expected: %.2f, but actual: %2.f", expected, actual)
		}
	})

	t.Run("Getting area of circle when radius is given using Circle struct's method", func(t *testing.T) {
		circle := Circle{1.0}
		actual := circle.Area()
		expected := math.Pi

		if actual != expected {
			t.Errorf("expected: %g, but actual: %g", expected, actual)
		}
	})
}
