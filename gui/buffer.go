package gui

import "github.com/jsniffen/geditor/edit"

type Buffer struct {
	X int
	Y int
	Width int
	Height int

	GapBuffer *edit.GapBuffer
}

func (b *Buffer) Render(cells []Cell, Width int) {
	x := b.X
	y := b.Y

	for _, c := range b.GapBuffer.String() {
		if c == ' ' {
			x += 1
			continue
		}

		if c == '\n' {
			x = b.X
			y += 1
			continue
		}

		cells[y*Width + x] = Cell{
			Color{255, 255, 255},
			Color{0, 0, 0},
			c,
		}

		x += 1
	}
}
