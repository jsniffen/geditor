package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jsniffen/geditor/gui"
	"github.com/jsniffen/geditor/term"
)

var Running = true
var Width uint32
var Height uint32

func main() {
	err := term.Init()
	if err != nil {
		log.Fatal(err)
	}

	Width, Height, err = term.GetSize()
	if err != nil {
		log.Fatal(err)
	}

	cells := make([]gui.Cell, Width*Height)
	for i := range cells {
		cells[i] = gui.Cell{
			gui.Color{255, uint8(i), 255},
			gui.Color{0, 0, 0},
			'X',
		}
	}

	chEvents := term.GetEvents()
	chTime := time.Tick(time.Second)

	for Running {
		select {
		case e := <-chEvents:
			if e.KeyCode == 'q' {
				Running = false
			}

		case _ = <-chTime:
		}

		term.Render(cells)
	}
}
