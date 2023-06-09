package gui

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Cell struct {
	Foreground Color
	Background Color
	Rune       rune
}

func StatusBar(w, h uint32, cells []Cell) {
	for i := 0; i < int(w); i += 1 {
		cells[int(w)*(int(h)-1) + i] = Cell{
			Color{255, 255, 255},
			Color{255, 255, 255},
			'X',
		}
	}
}
