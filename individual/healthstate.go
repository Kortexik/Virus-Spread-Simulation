package individual

type HealthState interface {
	HandleProximity(individual *Individual, other *Individual, dt float64)
	Update(individual *Individual, dt float64)
}
