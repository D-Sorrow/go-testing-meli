package simulator_test

import (
	"github.com/stretchr/testify/assert"
	"testdoubles/positioner"
	"testdoubles/simulator"
	"testing"
)

func TestCanCatch_Catch(t *testing.T) {
	positionerStub := positioner.PositionerStub{
		GetLinearDistanceFunc: func(from, to *positioner.Position) (linearDistance float64) {
			return 100
		}}
	method := simulator.NewCatchSimulatorDefault(100, &positionerStub)
	preyImp := simulator.Subject{Speed: 100, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
	huntImp := simulator.Subject{Speed: 120, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}

	outPut := method.CanCatch(&huntImp, &preyImp)
	assert.True(t, outPut)
}

func TestCanCatch_NotCatch(t *testing.T) {
	positionerStub := positioner.PositionerStub{
		GetLinearDistanceFunc: func(from, to *positioner.Position) (linearDistance float64) {
			return 100
		}}
	method := simulator.NewCatchSimulatorDefault(100, &positionerStub)
	preyImp := simulator.Subject{Speed: 130, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
	huntImp := simulator.Subject{Speed: 120, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}

	outPut := method.CanCatch(&huntImp, &preyImp)
	assert.False(t, outPut)
}

func TestCanCatch_NotCatchBySpeed(t *testing.T) {
	positionerStub := positioner.PositionerStub{
		GetLinearDistanceFunc: func(from, to *positioner.Position) (linearDistance float64) {
			return 10000
		}}
	method := simulator.NewCatchSimulatorDefault(100, &positionerStub)
	preyImp := simulator.Subject{Speed: 130, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
	huntImp := simulator.Subject{Speed: 140, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}

	outPut := method.CanCatch(&huntImp, &preyImp)
	assert.False(t, outPut)
}
