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

type Event struct {
	KeyCode byte
}
