package main

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

const backgroundColor = termbox.ColorBlue
const boardColor = termbox.ColorBlack
const instructionColor = termbox.ColorWhite

const defaultMarginWidth = 2
const defaultMarginHeight = 1
const titleStartX = defaultMarginWidth
const titleStartY = defaultMarginHeight
const titleHeight = 1
const titleEndY = titleStartY + titleHeight
const boardStartX = defaultMarginWidth
const boardStartY = titleEndY + defaultMarginHeight
const cellWidth = 2

const title = "Game"

func (g *Game) Render() {
	termbox.Clear(backgroundColor, backgroundColor)
	// Тақырып
	termboxPrint(titleStartX, titleStartY, instructionColor, backgroundColor, title)

	// Ойын тақтасы
	for i := 0; i < g.boardHeight; i++ {
		for j := 0; j < g.boardWidth; j++ {
			var cellColor termbox.Attribute

			// Food көрсету
			if i == g.FoodX && j == g.FoodY {
				cellColor = termbox.ColorRed
			} else if g.board[i][j] == 1 {
				cellColor = termbox.ColorGreen
			} else {
				cellColor = boardColor
			}

			for x := 0; x < cellWidth; x++ {
				termbox.SetCell(boardStartX+cellWidth*j+x, boardStartY+i, ' ', cellColor, cellColor)
			}
		}
	}

	// Score көрсету
	scoreMsg := "Score: " + strconv.Itoa(g.Score)
	termboxPrint(boardStartX, boardStartY+g.boardHeight+1, instructionColor, backgroundColor, scoreMsg)

	termbox.Flush()
}

func termboxPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
