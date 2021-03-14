package main

import (
	"log"
	"math"

	"github.com/patrickalexchan/introduction-to-go/3-variable/entity/grpc"
	c "github.com/patrickalexchan/introduction-to-go/3-variable/resource/car"
)

func main() {
	// Declare variable types in the beginning of function
	var (
		car    *c.Car // Custom Data Type
		status int    // Primitive Data Type
		err    error
	)

	// Example 1
	// Short Declaration
	// Usually use when inside the function for short used
	// if its used for the whole function and return, should be declared first
	for i := 0; i < 5; i++ {
		temp := math.Pow(float64(i), 2)
		log.Printf("Pow2 from %d = %.2f", i, temp)
	}

	// Example 2
	// Modify struct value with pointer
	// and should avoid catch data with _ variable

	// We dont want this direct modifier method
	// car = &c.Car{"H 3 LLO", "Black"}
	// car.Color = "Blue"
	// log.Printf("Current Car %+v \n", car)

	car = c.New("H 3 LLO", "Black")
	log.Printf("Current Car %+v \n", car)

	// _ = car.SetColor("White") --> Should be always catch the variable and validate when possible
	err = car.SetColor("White")
	if err != nil {
		log.Fatalln("Failed to set Car color")
	}
	log.Printf("Car after SetColor %+v \n", car)

	// Example 3
	// Const
	status = grpc.OK
	log.Println("GRPC Response Code : ", status)
}
