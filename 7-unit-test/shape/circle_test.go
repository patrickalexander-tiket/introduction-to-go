package shape

import (
	"testing"
)

func TestNewCircle(t *testing.T) {
	actualResult := NewCircle(1)
	if &actualResult == nil {
		t.Errorf("[TestNewCircle] failed to create new circle")
	}
}

func TestCircleArea(t *testing.T) {
	actualResult := NewCircle(2)
	if &actualResult == nil {
		t.Errorf("[TestNewCircle] failed to create new circle")
	}

	expectedResult := 12.566370614359172
	if actualResult.getArea() != expectedResult {
		t.Errorf("[TestCircleArea] failed to test circle.getArea()")
	}
}

func TestCircleAreaMultipleCase(t *testing.T) {
	type args struct {
		radius string
	}
	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{
			name:   "success",
			radius: 1,
			want:   3.141592653589793,
		},
		{
			name:   "success",
			radius: 2,
			want:   12.566370614359172,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := NewCircle(tt.radius)
			if got := actualResult.getArea(); got != tt.want {
				t.Errorf("keyOp.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
