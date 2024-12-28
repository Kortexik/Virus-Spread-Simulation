package main

import (
	. "TO/lab3/simulation"
	"TO/lab3/utility"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	os.RemoveAll("./saves")
	game := NewSimulation()
	ebiten.SetWindowSize(utility.ScreenWidth, utility.ScreenHeight)
	ebiten.SetWindowTitle("Symulacja")
	ebiten.SetTPS(25)

	// Run the simulation game
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
