package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)

	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	
	g := NewGame(20, 10)

	g.Render()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowDown:
					g.Direction = "Down"
					g.MoveDown()
				case ev.Key == termbox.KeyArrowLeft:
					g.Direction = "Left"
					g.MoveLeft()
				case ev.Key == termbox.KeyArrowRight:
					g.Direction = "Right"
					g.MoveRight()
				case ev.Key == termbox.KeyArrowUp:
					g.Direction = "Up"
					g.MoveUp()
				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC:
					return
				}
			}
		default:
			g.Render()
			time.Sleep(50 * time.Millisecond) // CPU жүктемесін азайту үшін
		}
	}
}
