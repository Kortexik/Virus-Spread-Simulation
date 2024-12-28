package individual

type ImmuneState struct{}

func (s *ImmuneState) HandleProximity(individual *Individual, other *Individual, dt float64) {
    // Immune individuals are not affected by infections.
}

func (s *ImmuneState) Update(individual *Individual, dt float64) {
    // No additional logic for immune individuals during update.
}
