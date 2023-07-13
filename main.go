package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
		vector.StrokeLine(screen, float32(i*cellSize), (cellSize / 6), float32(i*cellSize), (cellSize*3)-cellSize/6, float32(5), color.RGBA{50, 50, 50, 255}, true)
		vector.StrokeLine(screen, (cellSize / 6), float32(i*cellSize), (cellSize*3)-cellSize/6, float32(i*cellSize), float32(5), color.RGBA{50, 50, 50, 255}, true)
		for j := 0; j < 3; j++ {
			x, y := i*cellSize, j*cellSize
			switch board[i][j] {
			case 1:
				vector.StrokeLine(screen, float32(x)+(cellSize/3), float32(y)+(cellSize/3), float32(x)+(2*(cellSize/3)), float32(y)+(2*cellSize/3), float32(7), color.Black, true)
				vector.StrokeLine(screen, float32(x)+(2*(cellSize/3)), float32(y)+(cellSize/3), float32(x)+(cellSize/3), float32(y)+(2*(cellSize/3)), float32(7), color.Black, true)

			case -1:
				c := color.RGBA{255, 0, 0, 255} // Red for player 1
				vector.DrawFilledCircle(screen, float32(x+(cellSize/3)+(cellSize/6)), float32(y+(cellSize/3)+(cellSize/6)), (cellSize/3)-(cellSize/6), c, true)
				vector.DrawFilledCircle(screen, float32(x+(cellSize/3)+(cellSize/6)), float32(y+(cellSize/3)+(cellSize/6)), (cellSize / 7), color.White, true)
			default:
			}
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
