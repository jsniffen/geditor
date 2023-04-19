package main

import (
	"log"

	"github.com/jsniffen/geditor/term"
	"github.com/jsniffen/geditor/gui"
)

var Running = true
var Width = 0
var Height = 0

func main() {
	err := term.Init()
	if err != nil {
		log.Fatal(err)
	}

	cells := make([]gui.Cell, 1024)
	for i := range cells {
		cells[i] = gui.Cell{
			gui.Color{255, uint8(i), 255},
			gui.Color{0, 0, 0},
			'X',
		}
	}

	for Running {
		e, err := term.GetEvent()
		if err != nil {
			log.Fatal(err)
		}

		if e.KeyCode == 'q' {
			Running = false
		}

		term.Render(cells)
	}
}
