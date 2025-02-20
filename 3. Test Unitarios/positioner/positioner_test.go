package positioner

import (
	testify "github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPositioner_Negative(t *testing.T) {
	var from = Position{
		X: -100.0,
		Y: -100.0,
		Z: -100.0,
	}
	var to = Position{
		X: -150.0,
		Y: -150.0,
		Z: -150.0,
	}
	expectedValue := 86.60254037844386

	positionerStub := NewPositionerDefault()

	linear := positionerStub.GetLinearDistance(&from, &to)

	testify.Equal(t, expectedValue, linear)
}
func TestPositioner_Positive(t *testing.T) {
	var from = Position{
		X: 100.0,
		Y: 100.0,
		Z: 100.0,
	}
	var to = Position{
		X: 150.0,
		Y: 150.0,
		Z: 150.0,
	}
	expectedValue := 86.60254037844386

	positionerStub := NewPositionerDefault()

	linear := positionerStub.GetLinearDistance(&from, &to)

	testify.Equal(t, expectedValue, linear)
}

func TestPositioner_NotDecimal(t *testing.T) {
	var from = Position{
		X: 1.0,
		Y: 0.0,
		Z: 0.0,
	}
	var to = Position{
		X: 5.0,
		Y: 0.0,
		Z: 0.0,
	}
	positionerStub := NewPositionerDefault()

	linear := positionerStub.GetLinearDistance(&from, &to)

	isDecimal := linear != math.Floor(linear)

	testify.Equal(t, false, isDecimal)
}
