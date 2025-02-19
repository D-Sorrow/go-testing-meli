package prey

import "testdoubles/positioner"

type PreyStub struct {
	GetSpeedFunc    func() (speed float64)
	GetPositionFunc func() (position *positioner.Position)
}

func (ps PreyStub) NewPreyStub() *PreyStub {
	return &PreyStub{}
}

func (s *PreyStub) GetSpeed() float64 {
	return s.GetSpeedFunc()
}
func (s *PreyStub) GetPosition() *positioner.Position {
	return s.GetPositionFunc()
}
