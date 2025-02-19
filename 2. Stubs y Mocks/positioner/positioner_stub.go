package positioner

type PositionerStub struct {
	NegativeScene         bool
	PositiveScene         bool
	NotDecimalScene       bool
	GetLinearDistanceFunc func() (linearDistance float64)
}

func NewPositionerStub() *PositionerStub {
	return &PositionerStub{}
}

func (s PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	if s.NegativeScene || s.PositiveScene {
		return 87.0
	}
	return 87
}
