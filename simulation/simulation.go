package simulation

import (
	"TO/lab2/VectorsLib"
	. "TO/lab3/individual"
	. "TO/lab3/utility"
	"encoding/gob"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var inputStep string
var loadingError string

type Simulation struct {
	population  []*Individual
	paused      bool
	elapsedTime float64
	errorTime   time.Time
}

type SimulationMemento struct {
	Population  []ExportIndividual
	ElapsedTime float64
	Step        int
}

func (g *Simulation) SaveState(step int) error {
	memento := g.CreateMemento()
	filename := fmt.Sprintf("saves/step_%d", step)

	if err := os.MkdirAll("saves", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create saves directory: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create save file: %v", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(memento); err != nil {
		return fmt.Errorf("failed to encode memento: %v", err)
	}
	return nil
}

func (g *Simulation) LoadState(step int) error {
	filename := fmt.Sprintf("saves/step_%d", step)

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open save file: %v", err)
	}
	defer file.Close()

	memento := &SimulationMemento{}
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(memento); err != nil {
		return fmt.Errorf("failed to decode memento: %v", err)
	}

	g.ApplyMemento(memento)
	return nil
}

func (g *Simulation) CreateMemento() *SimulationMemento {
	idMap := make(map[*Individual]int)
	populationExport := make([]ExportIndividual, len(g.population))

	for i, ind := range g.population {
		idMap[ind] = i
		populationExport[i] = ind.ToExport(idMap)
	}

	return &SimulationMemento{
		Population:  populationExport,
		ElapsedTime: g.elapsedTime,
		Step:        currentStep,
	}
}

func (g *Simulation) ApplyMemento(memento *SimulationMemento) {
	idMap := make(map[int]*Individual)
	g.population = make([]*Individual, len(memento.Population))

	for i, expInd := range memento.Population {
		g.population[i] = expInd.ToIndividual(idMap)
		idMap[i] = g.population[i]
	}

	g.elapsedTime = memento.ElapsedTime
	currentStep = memento.Step
}

func (g *Simulation) DrawInputField(screen *ebiten.Image) {
	headerX := RectX + N + 20
	headerY := RectY + 400

	text.Draw(screen, "Enter Step to Load:", fontFace, int(headerX), int(headerY), color.White)
	text.Draw(screen, inputStep, fontFace, int(headerX), int(headerY+40), color.White)

	if loadingError != "" {
		if time.Since(g.errorTime).Seconds() < 1 { //nie czaje czemu nie > 1 w sensie ze uplynelo > 1 sekunde to wtedy mozna rysowac inny error lub ten sam znowu?
			text.Draw(screen, loadingError, fontFace, int(headerX), int(headerY+80), color.RGBA{255, 0, 0, 255})
		}
	}
}

func (g *Simulation) HandleInput() {
	runes := ebiten.AppendInputChars(nil)

	for _, char := range runes {
		inputStep += string(char)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(inputStep) > 0 {
		inputStep = inputStep[:len(inputStep)-1]
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && inputStep != "" {
		step, err := strconv.Atoi(inputStep)
		if err != nil {
			loadingError = "Invalid input! Please enter a valid number."
			g.errorTime = time.Now()
		} else {
			if loadErr := g.LoadState(step); loadErr != nil {
				loadingError = fmt.Sprintf("Failed to load step %d: %v", step, loadErr)
				g.errorTime = time.Now()
			} else {
				loadingError = ""
			}
		}
		inputStep = ""
	}
}

func (g *Simulation) GetPopulation() []*Individual {
	return g.population
}

func (g *Simulation) SetPopulation(newPopulation []*Individual) {
	g.population = newPopulation
}

var (
	fontFace    font.Face
	currentStep int
)

func init() {

	ttfData, err := os.ReadFile("../Doto-Regular.ttf")
	if err != nil {
		log.Fatalf("failed to read font file: %v", err)
	}

	tt, err := opentype.Parse(ttfData)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	const dpi = 72
	fontFace, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create font face: %v", err)
	}

	gob.Register(&HealthyState{})
	gob.Register(&InfectedState{})
	gob.Register(&ImmuneState{})
	gob.Register(&Individual{})
	gob.Register(&VectorsLib.Vector2D{})
}

func NewSimulation() *Simulation {
	individuals := make([]*Individual, I)
	for i := range individuals {
		individuals[i] = NewIndividual(
			RandomPostionInside(),
			RandomVelocity(),
			(&HealthyState{}),
			true,
		)
	}

	return &Simulation{population: individuals, paused: false}
}

func (g *Simulation) RemoveInactive() {

	var activeIndividuals []*Individual

	for _, individual := range g.population {
		if individual.Active {
			activeIndividuals = append(activeIndividuals, individual)
		}
	}

	g.population = activeIndividuals
}

func (g *Simulation) HandleInfecting(individual *Individual, dt float64) {
	for _, other := range g.GetPopulation() {
		if individual != other {
			individual.HandleProximity(other, dt)
		}
	}
}

func (g *Simulation) Update() error {
	dt := 1.0 / float64(ebiten.TPS())

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.paused = !g.paused
		return nil
	}

	if !g.paused {
		g.elapsedTime += dt
		currentStep++

		currentPopulation := g.GetPopulation()
		for _, individual := range currentPopulation {
			individual.Update(dt)
			g.HandleInfecting(individual, dt)

			if individual.CheckBounds(RectX, RectY, N, M) {
				switch rand.Intn(2) {
				case 0:
					individual.Active = false
				case 1:
					individual.MoveAway(RectX, RectY, N, M)
				}
			}

			if rand.Float64() < 0.001 {
				g.SetPopulation(append(currentPopulation, CreateIndividualOnBorder(RectX, RectY, N, M)))
			}
		}

		g.RemoveInactive()

		if err := g.SaveState(currentStep); err != nil {
			fmt.Println("Failed to save state:", err)
		}
	}

	g.HandleInput()
	return nil
}

func DrawLegend(screen *ebiten.Image, healthyCount int, infectedCount int, symptomaticCount int, immuneCount int, populationCount int, elapsedTime int, steps int) {
	headerText := fmt.Sprintf("Population Count: %d", populationCount)
	headerX := RectX + N + 20
	headerY := RectY - 40
	text.Draw(screen, headerText, fontFace, int(headerX), int(headerY), color.White)

	text.Draw(screen, fmt.Sprintf("Steps: %d", steps), fontFace, int(headerX), int(headerY)+220, color.White)
	text.Draw(screen, fmt.Sprintf("Seconds: %d", elapsedTime), fontFace, int(headerX)+250, int(headerY)+220, color.White)

	var legendX float32 = RectX + N + 20
	var legendY float32 = RectY
	var circleRadius float32 = 12
	var lineHeight float32 = 40

	vector.DrawFilledCircle(screen, legendX, legendY+circleRadius, circleRadius, Gray, false)
	text.Draw(screen, "Healthy: "+fmt.Sprint(healthyCount), fontFace, int(legendX+circleRadius*2+10), int(legendY+circleRadius+6), color.White)

	vector.DrawFilledCircle(screen, legendX, legendY+lineHeight+circleRadius, circleRadius, Orange, false)
	text.Draw(screen, "Infected (Asymptomatic): "+fmt.Sprint(infectedCount), fontFace, int(legendX+circleRadius*2+10), int(legendY+lineHeight+circleRadius+6), color.White)

	vector.DrawFilledCircle(screen, legendX, legendY+lineHeight*2+circleRadius, circleRadius, Red, false)
	text.Draw(screen, "Infected (Symptomatic): "+fmt.Sprint(symptomaticCount), fontFace, int(legendX+circleRadius*2+10), int(legendY+lineHeight*2+circleRadius+6), color.White)

	vector.DrawFilledCircle(screen, legendX, legendY+lineHeight*3+circleRadius, circleRadius, Green, false)
	text.Draw(screen, "Immune: "+fmt.Sprint(immuneCount), fontFace, int(legendX+circleRadius*2+10), int(legendY+lineHeight*3+circleRadius+6), color.White)
}

func (g *Simulation) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})
	vector.DrawFilledRect(screen, RectX, RectY, N, M, color.Black, false)

	healthyCount := 0
	infectedCount := 0
	symptomaticCount := 0
	immuneCount := 0

	for _, ind := range g.GetPopulation() {
		if ind.IsInfected() {
			if ind.HasSymptoms {
				symptomaticCount++
				vector.DrawFilledCircle(screen, float32(ind.GetX()), float32(ind.GetY()), 5, Red, false)
			} else {
				infectedCount++
				vector.DrawFilledCircle(screen, float32(ind.GetX()), float32(ind.GetY()), 5, Orange, false)
			}
		} else if ind.IsHealthy() {
			healthyCount++
			vector.DrawFilledCircle(screen, float32(ind.GetX()), float32(ind.GetY()), 5, Gray, false)
		} else if ind.IsImmune() {
			immuneCount++
			vector.DrawFilledCircle(screen, float32(ind.GetX()), float32(ind.GetY()), 5, Green, false)
		}
	}

	DrawLegend(screen, healthyCount, infectedCount, symptomaticCount, immuneCount, len(g.population), int(g.elapsedTime), currentStep)
	g.DrawInputField(screen)
}

func (g *Simulation) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
