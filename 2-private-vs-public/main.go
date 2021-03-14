package main

import "github.com/patrickalexchan/introduction-to-go/2-private-vs-public/resource/car"

func main() {
	car := car.Car{"H 3 LLO", "Black"}

	// Private Function
	// just can be called from internal package itself
	// carColor := car.color()

	// Public Function
	// can be called outside package
	car.GetColor()
}
