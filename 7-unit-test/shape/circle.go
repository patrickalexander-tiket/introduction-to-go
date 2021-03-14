package shape

import (
	"log"
	"math"
)

// Circle Struct
type Circle struct {
	Radius float64
}

// New creates new circle
func NewCircle(radius float64) Circle {
	return Circle{
		Radius: radius,
	}
}

func (c Circle) getArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

// PrintArea to print circle size
func (c Circle) PrintArea() {
	log.Printf("Circle Area : %f", c.getArea())
}
