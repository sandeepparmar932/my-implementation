package main

import "fmt"

type shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height * r.Width)
}

type Circle struct {
	Radius float64
}

// Implement the Shape interface for Circle
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func main() {
	rec := Rectangle{Width: 3.2, Height: 2}
	fmt.Println("Rectangle ", rec)
	fmt.Println(rec.Area())

	circ := Circle{Radius: 7}

	shapes := []shape{rec, circ}

	// Iterate over shapes and print their area and perimeter
	for _, shape := range shapes {
		fmt.Printf("Shape: %T, Area: %f, Perimeter: %f\n", shape, shape.Area(), shape.Perimeter())
	}
}
