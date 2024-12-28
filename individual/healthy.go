	package individual

	import (
		"TO/lab3/utility"
		"math/rand/v2"
	)

	type HealthyState struct{}

	func (s *HealthyState) HandleProximity(individual *Individual, other *Individual, dt float64) {
		if other.Active && other.IsInfected() && individual.GetPosition().CalculateDistance(other.GetPosition()) <= utility.InfectionRadius {
			individual.ProximityTimer[other] += dt
			if individual.ProximityTimer[other] >= 3.0 {
				infectionChance := 1.0
				if !other.HasSymptoms {
					infectionChance = 0.5
				}
				if rand.Float64() < infectionChance {
					individual.SetState(NewInfectedState(rand.Float64() < 0.5))
				}
			}
		} else {
			individual.ProximityTimer[other] = 0
		}
	}

	func (s *HealthyState) Update(individual *Individual, dt float64) {
		// No additional logic for healthy individuals during update.
	}
