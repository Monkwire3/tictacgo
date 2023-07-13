package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenSize = 500
	cellSize   = screenSize / 3
)

var (
	board = [3][3]int{}
	// players       = []string{"X", "O"}
	currentPlayer = 1
)

type Game struct{}

type Image struct {
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/cellSize, y/cellSize

		if board[i][j] == 0 {
			board[i][j] = currentPlayer
			currentPlayer *= -1
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for i := 0; i < 3; i++ {
		ebitenutil.DrawLine(screen, float64(i*cellSize), (cellSize / 6), float64(i*cellSize), (cellSize*3)-cellSize/6, color.RGBA{50, 50, 50, 255})
		ebitenutil.DrawLine(screen, (cellSize / 6), float64(i*cellSize), (cellSize*3)-cellSize/6, float64(i*cellSize), color.RGBA{50, 50, 50, 255})
		for j := 0; j < 3; j++ {
			x, y := i*cellSize, j*cellSize
			var c color.Color
			switch board[i][j] {
			case 1:
				c = color.RGBA{255, 0, 0, 255} // Red for player 1
			case -1:
				c = color.RGBA{0, 0, 255, 255} // Blue for player 2
			default:
				c = color.White
			}
			ebitenutil.DrawRect(screen, float64(x)+(cellSize/3), float64(y)+(cellSize/3), cellSize/3, cellSize/3, c)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenSize, screenSize
}

func main() {
	ebiten.SetWindowSize(screenSize, screenSize)
	ebiten.SetWindowTitle("TicTacGo")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
