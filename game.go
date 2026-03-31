package main

import (
	"math/rand"
	"time"
)

type Game struct {
	board       [][]int
	pX          int
	pY          int
	boardHeight int
	boardWidth  int
	Direction   string
	Score       int
	FoodX       int
	FoodY       int
}

// NewGame енді параметр қабылдайды
func NewGame(width, height int) *Game {
	g := &Game{
		boardHeight: height,
		boardWidth:  width,
	}
	g.ResetGame()
	rand.Seed(time.Now().UnixNano())
	g.PlaceFood()
	return g
}

// Тақтаны бастапқы күйге келтіру
func (g *Game) ResetGame() {
	g.board = make([][]int, g.boardHeight)
	for i := 0; i < g.boardHeight; i++ {
		g.board[i] = make([]int, g.boardWidth)
		for j := 0; j < g.boardWidth; j++ {
			g.board[i][j] = 0
		}
	}
	g.board[0][0] = 1
	g.pX = 0
	g.pY = 0
	g.Score = 0
}

// Қозғалыс функциялары
func (g *Game) MoveUp() {
	g.tryMove(-1, 0)
}

func (g *Game) MoveDown() {
	g.tryMove(1, 0)
}

func (g *Game) MoveLeft() {
	g.tryMove(0, -1)
}

func (g *Game) MoveRight() {
	g.tryMove(0, 1)
}

// Жалпы қозғалыс функциясы
func (g *Game) tryMove(dx, dy int) {
	newX := g.pX + dx
	newY := g.pY + dy

	// шекара және өзіне соғу
	if newX < 0 || newX >= g.boardHeight || newY < 0 || newY >= g.boardWidth {
		return
	}

	// Жаңарту
	g.board[g.pX][g.pY] = 0
	g.pX = newX
	g.pY = newY
	g.board[g.pX][g.pY] = 1

	// Food тексеру
	if g.pX == g.FoodX && g.pY == g.FoodY {
		g.Score++
		g.PlaceFood()
	}
}

// Food орналастыру
func (g *Game) PlaceFood() {
	for {
		x := rand.Intn(g.boardHeight)
		y := rand.Intn(g.boardWidth)
		if g.board[x][y] == 0 {
			g.FoodX = x
			g.FoodY = y
			break
		}
	}
}

// Дирекция тексеру (main.go үшін қажет болса)
func (g *Game) CheckMove() bool {
	switch g.Direction {
	case "Up":
		return g.pX > 0 && g.board[g.pX-1][g.pY] != 1
	case "Right":
		return g.pY < g.boardWidth-1 && g.board[g.pX][g.pY+1] != 1
	case "Down":
		return g.pX < g.boardHeight-1 && g.board[g.pX+1][g.pY] != 1
	case "Left":
		return g.pY > 0 && g.board[g.pX][g.pY-1] != 1
	default:
		return false
	}
}
