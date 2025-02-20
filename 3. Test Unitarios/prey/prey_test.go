package prey

import (
	"github.com/stretchr/testify/assert"
	"testdoubles/positioner"
	"testing"
)

func TestPrey_Default(t *testing.T) {
	tunaDefault := NewTuna(0, nil)

	speedExpected := 0
	var positionExpected *positioner.Position = nil

	assert.Equal(t, speedExpected, tunaDefault.GetSpeed())
	assert.Equal(t, positionExpected, tunaDefault.GetPosition())
}

func TestPrey_Different(t *testing.T) {
	tunaDefault := NewTuna(252.0, &positioner.Position{X: 0.0, Y: 0.0, Z: 1.0})

	speedExpected := 252.0
	positionExpected := &positioner.Position{
		X: 0.0,
		Y: 0.0,
		Z: 1.0,
	}

	assert.Equal(t, speedExpected, tunaDefault.GetSpeed())
	assert.Equal(t, positionExpected, tunaDefault.GetPosition())
}
