package structs

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "perimeter of a rectangle", shape: Rectangle{10.0, 10.0}, expected: 40.0},
		{name: "perimeter of a circle", shape: Circle{10.0}, expected: 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.expected {
				t.Errorf("got %g, expected %g", got, tt.expected)
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "area of a rectangle", shape: Rectangle{12, 6}, expected: 72.0},
		{name: "area of a circle", shape: Circle{10}, expected: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expected {
				t.Errorf("got %g, expected %g", got, tt.expected)
			}
		})
	}

}
