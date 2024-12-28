package utility

import (
	"TO/lab2/VectorsLib"
	"image/color"
	"math/rand/v2"
)

const (
	ScreenWidth     = 1800
	ScreenHeight    = 980
	N               = 800
	M               = 800
	RectX           = ScreenWidth/2 - N/2
	RectY           = ScreenHeight/2 - M/2
	CenterX         = RectX + N/2
	CenterY         = RectY + M/2
	I               = 100
	InfectionRadius = 50
)

var (
	Red    = color.RGBA{236, 78, 50, 1}
	Green  = color.RGBA{59, 170, 86, 1}
	Gray   = color.RGBA{109, 95, 95, 1}
	Orange = color.RGBA{255, 155, 1, 1}
)

func RandomBorderPosition() *VectorsLib.Vector2D {
	var randBorderX, randBorderY float64

	switch rand.IntN(4) {
	case 0: // Top border
		randBorderX = rand.Float64()*N + RectX
		randBorderY = RectY
	case 1: // Bottom border
		randBorderX = rand.Float64()*N + RectX
		randBorderY = RectY + M
	case 2: // Left border
		randBorderX = RectX
		randBorderY = rand.Float64()*M + RectY
	case 3: // Right border
		randBorderX = RectX + N
		randBorderY = rand.Float64()*M + RectY
	}

	return VectorsLib.NewVector2D(randBorderX, randBorderY)
}

func RandomPostionInside() *VectorsLib.Vector2D {
	randInsideX := rand.Float64()*N + RectX
	randInsideY := rand.Float64()*M + RectY

	return VectorsLib.NewVector2D(randInsideX, randInsideY)
}

func RandomVelocity() *VectorsLib.Vector2D {
	return VectorsLib.NewVector2D(rand.Float64()*5-2.5, rand.Float64()*5-2.5)
}

func RandomTenPercent() bool {
	switch rand.IntN(10) {
	case 9:
		return true
	default:
		return false
	}
}

func RandomFiftyPercent() bool {
	switch rand.IntN(2) {
	case 0:
		return true
	default:
		return false
	}
}
