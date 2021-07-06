// Define an interface which defines a method area().  Create types for square,
// rectangle and circle, and ensure they satisfy your interface.  Create a
// function that accepts a value of your interface type and outputs the area,
// and call this function for different shapes.
//
// Template available at: http://play.golang.org/p/rL5tT2VTJH
package main

// Add your imports here.
import (
	"math"
	"fmt"
)

// Define an interface with a method area().  Make sure you use a meaningful
// name and a sensible return type.

type areaer interface {
	area() float32
}

// Create square, rectangle and circle types, and ensure they satisfy your
// interface (you'll need to use the `Pi` constant from the `math` package for
// your circle).

type square struct {
	side float32
}

func (s square) area() float32 {
	return s.side * s.side
}

type rectangle struct {
	width float32
	height float32
}

func (r rectangle) area() float32 {
	return r.width * r.height
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

// Write a function that accepts a value of your interface and outputs the
// area.

func outputArea(shape areaer) {
	fmt.Printf("%T with area %f\n", shape, shape.area())
}

func main() {
	// Create a slice of your interface type, and populate it with a number of
	// different shapes.

	shapes := []areaer{
		square{4.3},
		rectangle{2, 3.4},
		circle{2.2},
	}

	// Loop through your shapes and use your function to output the area of
	// each one.

	for _, shape := range shapes {
		outputArea(shape)
	}
}
