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

// func (g *Game) Update() error {
// 	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
// 		x, y := ebiten.CursorPosition()
// 		i, j := x/cellSize, y/cellSize

//			if board[i][j] == 0 {
//				board[i][j] = players[currentPlayer]
//				if currentPlayer == 0 {
//					currentPlayer++
//				} else {
//					currentPlayer--
//
//			}
//			return nil
//		}
//	}
func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/cellSize, y/cellSize

		if board[i][j] == 0 {
			board[i][j] = currentPlayer
		}

		currentPlayer *= -1

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			x, y := i*cellSize, j*cellSize
			var c color.Color
			switch board[i][j] {
			case 1:
				c = color.RGBA{255, 0, 0, 255} // Red for player 1
			case -1:
				c = color.RGBA{0, 0, 255, 255} // Blue for player 2
			default:
				c = color.RGBA{0, 0, 0, 255} // Black for empty cells
			}
			ebitenutil.DrawRect(screen, float64(x), float64(y), cellSize, cellSize, c)
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
