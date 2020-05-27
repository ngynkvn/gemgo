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
	screenX     int
	screenY     int
}

func (w Window) DrawLine(y int, line gemini.Line, maxWidth int) {
	for x, chr := range line.Raw {
		w.screen.SetContent(x, y, chr, nil, tcell.StyleDefault)
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
	if w.screenY > 0 {
		w.screenY--
	}
}

func (w *Window) ScrollDown() {
	if w.screenY <= len(w.contents.Lines) {
		w.screenY++
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
	if key, ok := event.(*tcell.EventKey); ok {
		w.status = fmt.Sprintf("key %s", key.Name())
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

func (w *Window) Render() {
	width, height := w.screen.Size()
	body := w.contents
	w.screen.Clear()
	for y := 0; y < height && (y < len(body.Lines)-w.screenY); y++ {
		line := body.Lines[y+w.screenY]
		w.DrawLine(y, line, width)
	}
	w.RenderStatusBar()
	w.screen.Show()
}

func NewWindow() Window {
	screen, err := tcell.NewScreen()
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
