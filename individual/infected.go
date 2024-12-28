package individual

import (
	"math/rand"
)

type InfectedState struct {
	hasSymptoms       bool
	TimeSpentInfected float64
}

func NewInfectedState(hasSymptoms bool) *InfectedState {
	return &InfectedState{
		hasSymptoms:       hasSymptoms,
		TimeSpentInfected: 0,
	}
}

func (s *InfectedState) HandleProximity(individual *Individual, other *Individual, dt float64) {
	// Infected individuals don't infect themselves further.
}

func (s *InfectedState) Update(individual *Individual, dt float64) {
	s.TimeSpentInfected += dt
	if s.TimeSpentInfected > float64(20+rand.Intn(11)) {
		individual.SetState(&ImmuneState{})
	}
}
