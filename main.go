package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenSize = 500
	cellSize   = screenSize / 3
)

var (
	board         = [3][3]int{}
	currentPlayer = 1
)

type Game struct {
	gameWon bool
	winner  int
	inGame  bool
	inMenu  bool
	inEnd   bool
}

type Image struct {
}

func (g *Game) StartGame() {
	g.gameWon = false
	g.inGame = true
	g.winner = 0
	g.inMenu = false
	g.inEnd = false
}

func (g *Game) EndGame(winner int) {
	g.inGame = false
	g.gameWon = true
	g.winner = winner
	g.inMenu = false
	g.inEnd = true
}

func (g *Game) ReturnToMenu() {
	g.inGame = false
	g.gameWon = false
	g.winner = 0
	g.inMenu = true
	g.inEnd = false
}

func (g *Game) Update() error {
	if g.inGame {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			i, j := x/cellSize, y/cellSize

			if board[i][j] == 0 {
				board[i][j] = currentPlayer
				currentPlayer *= -1
			}
		}
	} else if g.inMenu {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			g.StartGame()
		}
	} else if g.inEnd {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			g.ReturnToMenu()
		}
	} else {
		g.ReturnToMenu()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.inGame {
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
	} else if g.inMenu {
		screen.Fill(color.Black)
		vector.StrokeRect(screen, float32((screenSize/2)-(screenSize/4)), float32((screenSize/2)-(screenSize/4)), float32(screenSize/2), float32(screenSize/4), float32(5), color.White, false)
		ebitenutil.DebugPrintAt(screen, "TicTacGo\n Click to play.", (screenSize/2)-(screenSize/8), (screenSize/2)-(screenSize/6))

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
