package main

import (
	"log"
	"math"

	"github.com/patrickalexchan/introduction-to-go/7-unit-test/entity"
	"github.com/patrickalexchan/introduction-to-go/7-unit-test/shape"
)

func main() {
	var (
		circ   shape.Circle // Custom Data Type
		status int          // Primitive Data Type
	)

	circ = shape.Circle{20}

	// Private Function
	circ.PrintArea()

	// Const
	status = entity.DEADLINE_EXCEEDED
	log.Println("Response Code : ", status)

	// Short Declaration
	// Usually use when inside the function for short used
	// if its used for the whole function and return, should be declared first
	for i := 0; i < 5; i++ {
		temp := math.Pow(float64(i), 2)
		log.Printf("Pow2 from %d = %.2f", i, temp)
	}
}
