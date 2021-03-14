package main

import (
	"log"
	"math"

	"github.com/patrickalexchan/introduction-to-go/3-variable/entity"
	"github.com/patrickalexchan/introduction-to-go/3-variable/shape"
)

func main() {
	// Declare variable types in the beginning of function
	var (
		rect     shape.Rectangle // Custom Data Type
		circ     shape.Circle    // Custom Data Type
		rectArea float64         // Primitive Data Type
		status   int             // Primitive Data Type
	)

	// Example 1
	// Const
	status = entity.DEADLINE_EXCEEDED
	log.Println("Response Code : ", status)

	// Example 2
	// Short Declaration
	// Usually use when inside the function for short used
	// if its used for the whole function and return, should be declared first
	for i := 0; i < 5; i++ {
		temp := math.Pow(float64(i), 2)
		log.Printf("Pow2 from %d = %.2f", i, temp)
	}

	rect = shape.Rectangle{Width: 7, Height: 8}
	circ = shape.Circle{20}

	// Public Function
	rectArea = rect.Area()
	log.Printf("Rectangle Area : %f \n", rectArea)

	// Private Function
	circ.PrintArea()

}
