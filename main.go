package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const width = 1920
const height = 1080

type App struct{}

func (a *App) Update() error {
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	file, err := os.OpenFile("out.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println("start app")
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("my-project")
	if err := ebiten.RunGame(&App{}); err != nil {
		log.Fatal(err)
	}
}
