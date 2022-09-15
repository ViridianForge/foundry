package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// Perimeter implementations
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t Triangle) Perimeter() float64 {
	return t.Width + t.Height + math.Sqrt((t.Width*t.Width)+(t.Height*t.Height))
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area implementations
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return t.Width * t.Height * 0.5
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
