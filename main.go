package main

import (
	"log"
	"time"

	"github.com/jsniffen/geditor/gui"
	"github.com/jsniffen/geditor/term"
	"github.com/jsniffen/geditor/edit"
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

	cells := make([]gui.Cell, int(Width*Height))
	for i := range cells {
		cells[i] = gui.Cell{
			gui.Color{0, 0, 0},
			gui.Color{0, 0, 0},
			' ',
		}
	}

	chEvents := term.GetEvents()
	chTime := time.Tick(time.Second)
	gb := edit.NewGapBuffer(100)
	gb.Insert('h');
	gb.Insert('e');
	gb.Insert('l');
	gb.Insert('l');
	gb.Insert('o');

	b := gui.Buffer{0, 0, 100, 100, gb}

	for Running {
		for i := range cells {
			cells[i] = gui.Cell{
				gui.Color{0, 0, 0},
				gui.Color{0, 0, 0},
				' ',
			}
		}

		b.Render(cells, int(Width))

		term.Render(cells)

		select {
		case e := <-chEvents:
			if e.KeyCode == 'q' {
				Running = false
			} else if e.KeyCode == term.KeyDelete {
				gb.Delete()
			} else {
				gb.Insert(e.KeyCode)
			}


		case _ = <-chTime:
		}
	}
}
