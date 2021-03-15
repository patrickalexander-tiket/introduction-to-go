package car

import (
	"testing"
)

func TestNewCar(t *testing.T) {
	actualResult := New("H 3 LLO", "White")
	if actualResult == nil {
		t.Errorf("[TestNewCar] failed to create new car")
	}
}
func TestGetColor(t *testing.T) {
	actualResult := New("H 3 LLO", "White")
	expectedResult := "White"
	if actualResult.GetColor() != "White" {
		t.Errorf("[TestGetColor] Expected %s but got %v", expectedResult, nil)
	}
}

func TestGetPlateNum(t *testing.T) {
	actualResult := New("H 3 LLO", "White")
	expectedResult := "H 3 LLO"
	if actualResult.GetPlateNum() != expectedResult {
		t.Errorf("[TestGetPlateNum] Expected %s but got %v", expectedResult, nil)
	}
}

func TestGetCarMultipleCase(t *testing.T) {
	tests := []struct {
		name        string
		plateNumber string
		color       string
		wantPlate   string
		wantColor   string
	}{
		{
			name:        "success",
			plateNumber: "H 3 LLO",
			color:       "White",
			wantPlate:   "H 3 LLO",
			wantColor:   "White",
		},
		{
			name:        "success",
			plateNumber: "H 4 LLO",
			color:       "Black",
			wantPlate:   "H 4 LLO",
			wantColor:   "Black",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := New(tt.plateNumber, tt.color)
			if got := actualResult.GetColor(); got != tt.wantColor {
				t.Errorf("car.GetColor() = %v, want %v", got, tt.wantColor)
			}
			if got := actualResult.GetPlateNum(); got != tt.wantPlate {
				t.Errorf("car.GetPlateNum() = %v, want %v", got, tt.wantPlate)
			}
		})
	}
}
