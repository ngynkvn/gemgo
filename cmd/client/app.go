package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/ngynkvn/gemgo/gemini"
)

type Window struct {
	screen      tcell.Screen
	status      string
	contents    gemini.Body
	shouldClose bool
	offsetX     int
	offsetY     int
	cursorX     int
	cursorY     int
}

func (w Window) DrawLine(y int, line gemini.Line, maxWidth int) {
	display, lineStyle := line.Display()
	style := tcell.StyleDefault.Underline(lineStyle == gemini.LinkStyle)

	for x, chr := range display {
		w.screen.SetContent(x, y, chr, nil, style)
	}
	// switch line.LineType {
	// case gemini.Link:
	// default:
	// }
}

func (w Window) RenderStatusBar() {
	width, height := w.screen.Size()
	style := tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tcell.ColorBlack)
	for x := 0; x < width; x++ {
		w.screen.SetContent(x, height-1, ' ', nil, style)
	}
	for x, chr := range w.status {
		w.screen.SetContent(x, height-1, chr, nil, style)
	}
}

func (w *Window) Close() {
	w.screen.Clear()
	w.screen.Sync()
	w.shouldClose = true
}

func (w *Window) ScrollUp() {
	if w.offsetY > 0 {
		w.offsetY--
	}
}

func (w *Window) ScrollDown() {
	if w.offsetY < len(w.contents.Lines)-1 {
		w.offsetY++
	}
}

func (w *Window) handleEnter() {
	// !! TODO !! This is gonna introduce a bug later
	line := w.contents.Lines[w.offsetY+w.cursorY]
	if link, ok := line.(gemini.Link); ok {
		w.status = fmt.Sprintf("GO(%s)", link.Link())
	}
}

func (w *Window) handleEvent(event tcell.Event) {
	if err, ok := event.(*tcell.EventError); ok {
		w.status = err.Error()
	}
	if _, ok := event.(*tcell.EventInterrupt); ok {
		//TODO
		w.status = "interrupt"
	}
	if mouse, ok := event.(*tcell.EventMouse); ok {
		x, y := mouse.Position()
		w.status = fmt.Sprintf("MouseEvt(%d, %d)", x, y)
	}
	if key, ok := event.(*tcell.EventKey); ok {
		w.status = fmt.Sprintf("Key(%s)", key.Name())
		switch key.Key() {
		case tcell.KeyCtrlC:
			w.Close()
		case tcell.KeyRune:
			if key.Rune() == 'q' {
				w.Close()
			}
			if key.Rune() == 'j' {
				w.ScrollDown()
			}
			if key.Rune() == 'k' {
				w.ScrollUp()
			}
		case tcell.KeyDown:
			w.ScrollDown()
		case tcell.KeyUp:
			w.ScrollUp()
		case tcell.KeyEnter:
			w.handleEnter()
		}
	}
}

func (w *Window) Run() {
	for !w.shouldClose {
		w.Render()
		evt := w.screen.PollEvent()
		w.handleEvent(evt)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (w *Window) Render() {
	width, height := w.screen.Size()
	body := w.contents
	w.screen.Clear()
	w.screen.ShowCursor(w.cursorX, w.cursorY)
	for y := 0; y < height; y++ {
		if y+w.offsetY < len(body.Lines) {
			line := body.Lines[y+w.offsetY]
			w.DrawLine(y, line, width)
		}
	}
	w.RenderStatusBar()
	w.screen.Show()
}

func NewWindow() Window {
	screen, err := tcell.NewScreen()
	screen.EnableMouse()
	if err != nil {
		log.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}
	var window Window
	window.screen = screen
	return window
}

func main() {
	args := os.Args
	url := gemini.ParseURL(args[1])
	log.Println("Visiting", url.String())

	gc := gemini.NewGeminiConnection(url)

	gc.SendRequest(url)
	header := gc.ReceiveHeader()
	fmt.Println("header", header)
	body := gc.ReceiveBody()
	window := NewWindow()
	window.contents = body
	window.status = "Hi Kevin"
	window.Run()
	// w, h := screen.Size()
	// log.Printf("Screen (W,H) : (%d,%d)\n", w, h)
}
