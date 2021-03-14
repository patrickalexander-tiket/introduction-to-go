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

func (c Car) color() string {
	return c.Color
}

// GetColor to get the color of the car
func (c Car) GetColor() string {
	return c.color()
}

// SetColor the color of the car
func (c *Car) SetColor(color string) error {
	if color == "" {
		return errors.New("Must specify car color")
	}

	c.Color = color

	return nil
}

// GetPlateNum to get plate number of the car
func (c Car) GetPlateNum() string {
	return c.PlateNumber
}
