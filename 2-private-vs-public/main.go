package main

import (
	"log"

	"github.com/patrickalexchan/introduction-to-go/2-private-vs-public/shape"
)

func main() {
	rect := shape.Rectangle{Width: 7, Height: 8}
	circ := shape.Circle{20}

	// Public Function
	rectArea := rect.Area()
	log.Printf("Rectangle Area : %f \n", rectArea)

	// Private Function
	circ.PrintArea()

}
