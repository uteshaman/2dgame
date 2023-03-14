package main

import "github.com/nsf/termbox-go"

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

	termGame.NewGame()

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
				case ev.ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlc:
					return
				}
			}
		default:
			g.Render()
		}
	}

}
