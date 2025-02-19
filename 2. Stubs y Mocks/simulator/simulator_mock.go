package simulator

type SimulatorMock struct {
	CanCatchFunc func(hunter, prey *Subject) (canCatch bool)
	Count        int
}

func NewSimulatorMock() *SimulatorMock {
	return &SimulatorMock{}
}

func (sm SimulatorMock) CanCatch(hunter, prey *Subject) (canCatch bool) {
	sm.Count++
	return sm.CanCatchFunc(hunter, prey)
}
