package car

import "errors"

// Car Class contain of PlateNumber and color
type Car struct {
	PlateNumber string
	Color       string
}

// New creates new Car, dont need reference since the car will not modified
func New(plateNumber string, color string) *Car {
	return &Car{
		PlateNumber: plateNumber,
		Color:       color,
	}
}

//START OMIT
func (c Car) color() string { // HL12
	return c.Color
} // HL12

// GetColor to get the color of the car
func (c Car) GetColor() string { // HL12
	return c.Color
} // HL12

//END OMIT

// GetColor to get the color of the car
func (c *Car) SetColor(color string) (car Car, err error) {
	if color == "" {
		return car, errors.New("Must specify car color")
	}

	c.Color = color

	return car, nil
}

// GetPlateNum to get plate number of the car
func (c Car) GetPlateNum() string {
	return c.PlateNumber
}
