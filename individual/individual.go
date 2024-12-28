package individual

import (
	vectors "TO/lab2/VectorsLib"
	"TO/lab3/utility"
	"fmt"
	"math/rand"
)

type Individual struct {
	position       *vectors.Vector2D
	velocity       *vectors.Vector2D
	state          HealthState
	Active         bool
	ProximityTimer map[*Individual]float64
	HasSymptoms    bool
}

type ExportIndividual struct {
	Position       [2]float64
	Velocity       [2]float64
	State          string
	Active         bool
	HasSymptoms    bool
	ProximityTimer map[int]float64
}

func (ind *Individual) ToExport(idMap map[*Individual]int) ExportIndividual {
	proximityTimer := make(map[int]float64)
	for other, timer := range ind.ProximityTimer {
		proximityTimer[idMap[other]] = timer
	}

	return ExportIndividual{
		Position:       [2]float64{ind.GetX(), ind.GetY()},
		Velocity:       [2]float64{ind.GetVelocity().GetX(), ind.GetVelocity().GetY()},
		State:          fmt.Sprintf("%T", ind.state),
		Active:         ind.Active,
		HasSymptoms:    ind.HasSymptoms,
		ProximityTimer: proximityTimer,
	}
}

func (exp *ExportIndividual) ToIndividual(idMap map[int]*Individual) *Individual {
	pos := vectors.NewVector2D(exp.Position[0], exp.Position[1])
	vel := vectors.NewVector2D(exp.Velocity[0], exp.Velocity[1])

	var state HealthState
	switch exp.State {
	case "*individual.HealthyState":
		state = &HealthyState{}
	case "*individual.InfectedState":
		state = &InfectedState{hasSymptoms: exp.HasSymptoms}
	case "*individual.ImmuneState":
		state = &ImmuneState{}
	}

	ind := NewIndividual(pos, vel, state, exp.Active)
	for id, timer := range exp.ProximityTimer {
		ind.ProximityTimer[idMap[id]] = timer
	}

	return ind
}

func (individual *Individual) SetState(state HealthState) {
	individual.state = state
	if infectedState, ok := state.(*InfectedState); ok {
		individual.HasSymptoms = infectedState.hasSymptoms
	} else {
		individual.HasSymptoms = false
	}
}

func (individual *Individual) GetState() HealthState {
	return individual.state
}

func (individual *Individual) HandleProximity(other *Individual, dt float64) {
	individual.state.HandleProximity(individual, other, dt)
}

func (individual *Individual) Update(dt float64) {
	individual.state.Update(individual, dt)
	individual.Move(dt)
}

func (individual *Individual) GetPosition() *vectors.Vector2D {
	return individual.position
}

func (individual *Individual) GetVelocity() *vectors.Vector2D {
	return individual.velocity
}
func (individual *Individual) SetPosition(newPosition *vectors.Vector2D) {
	individual.position = newPosition
}

func (individual *Individual) SetVelocity(newVelocity *vectors.Vector2D) {
	individual.velocity = newVelocity
}

func (individual *Individual) GetX() float64 {
	return individual.position.GetX()
}

func (individual *Individual) GetY() float64 {
	return individual.position.GetY()
}

func NewIndividual(position *vectors.Vector2D, velocity *vectors.Vector2D, state HealthState, active bool) *Individual {
	return &Individual{
		position:       position,
		velocity:       velocity,
		state:          state,
		Active:         active,
		ProximityTimer: make(map[*Individual]float64),
	}
}

func (individual *Individual) Move(dt float64) {

	dx := individual.velocity.GetX() * dt
	dy := individual.velocity.GetY() * dt

	individual.position.Set(
		individual.position.GetX()+dx,
		individual.position.GetY()+dy,
	)
}

func (individual *Individual) CheckBounds(rectX float64, rectY float64, N float64, M float64) bool {
	if individual.position.GetX() <= rectX || individual.position.GetX() >= rectX+N ||
		individual.position.GetY() <= rectY || individual.position.GetY() >= rectY+M {
		return true
	}
	return false
}

func (individual *Individual) MoveAway(rectX float64, rectY float64, N float64, M float64) {
	centerX, centerY := rectX+N/2, rectY+M/2

	if individual.position.GetX() <= rectX {
		individual.position.SetX(rectX + 2.5)
	} else if individual.position.GetX() >= rectX+N {
		individual.position.SetX(rectX + N - 2.5)
	}

	if individual.position.GetY() <= rectY {
		individual.position.SetY(rectY + 2.5)
	} else if individual.position.GetY() >= rectY+M {
		individual.position.SetY(rectY + M - 2.5)
	}

	newVelocity := vectors.NewVector2D(
		centerX-individual.position.GetX(),
		centerY-individual.position.GetY(),
	)
	newVelocity.Normalize()
	newVelocity.Scale(2.5)
	individual.SetVelocity(newVelocity)
}

func CreateIndividualOnBorder(rectX float64, rectY float64, N float64, M float64) *Individual {
	borderPosition := utility.RandomBorderPosition()

	newIndividual := NewIndividual(
		vectors.NewVector2D(borderPosition.GetX(), borderPosition.GetY()),
		utility.RandomVelocity(),
		&HealthyState{},
		true,
	)

	if utility.RandomTenPercent() {
		infectedState := NewInfectedState(rand.Float64() < 0.5)
		newIndividual.SetState(infectedState)
	}

	return newIndividual

}

func (ind *Individual) IsInfected() bool {
	_, ok := ind.state.(*InfectedState)
	return ok
}

func (ind *Individual) IsImmune() bool {
	_, ok := ind.state.(*ImmuneState)
	return ok
}

func (ind *Individual) IsHealthy() bool {
	_, ok := ind.state.(*HealthyState)
	return ok
}
