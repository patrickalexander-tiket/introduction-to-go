package shape

import (
	"log"
	"math"
)

// Circle Struct
type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// PrintArea to print circle size
func (c Circle) PrintArea() {
	log.Printf("Circle Area : %f", c.area())
}
