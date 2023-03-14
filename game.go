package main

type Game struct {
	board       [][]int
	pX          int
	pY          int
	boardHeight int
	boardWidth  int
	Direction string 
}

func NewGame() *Game {
	g := new(Game)

	g.boardHeight = 16
	g.boardWidth = 16

	g.ResetGame()

	return g

}
func (g *Game) resetGame() {
	g.board = make([][]int, g.boardHeight)

	for i := 0; i < g.boardHeight; i++ {
		g.board[i] = make([][]int, g.boardWidth)

		for j := 0; j < g.boardHeight; j++ {
			g.board[i][j] = 0
		}
	}
	g.board[0][0] = 1
	g.pX = 0
	g.pY = 0

}
func (g *Game) MoveUp() {
	g.board[g.pX][g.pY] = 0
	g.pX--
	g.board[g.pX][g.pY] = 1
}
func (g *Game) MoveRight() {
	g.board[g.pX][g.pY] = 0
	g.pY++
	g.board[g.pX][g.pY] = 1
}
func (g *Game) MoveDown() {
	g.board[g.pX][g.pY] = 0
	g.pX++
	g.board[g.pX][g.pY] = 1
}
func (g *Game) MoveLeft() {
	g.board[g.pX][g.pY] = 0
	g.pY--
	g.board[g.pX][g.pY] = 1
}

for (g *Game) CheckMOve() bool{
	switch g.Direction{

	case "up":
		if g.pX == 0 || g.board[g.pX-1][g.pY] == 1{
			return false
		}
		return true 
		
	case "right":
		if g.pY == g.boardHeight-1 || g.board[g.pX][g.pY+1] == 1{
			return false
		}
		return true

	case "down":
		if g.pX == g.boardWidth-1 || g.board[g.pX+1][g.pY] == 1{
			return false
		}
		return true 

	case "left":
		if g.pY == 0 || g.board[g.pX][g.pY-1] == 1{
			return false
		}
		return true
	default:
		panic("OMG!!! Direction is wrong")
	
	}
}
