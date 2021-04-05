package main

import (
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
	"log"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Hello World!"
	p.SetRect(0, 0, 25, 5)

	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			p.Text += " " + e.ID
			ui.Render(p)
			//break
		}
	}
}