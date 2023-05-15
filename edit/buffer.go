package edit

type Buffer struct {
	path       string
	pieceTable *PieceTable
	cursorX    int
	cursorY    int
	x          int
	y          int
	width      int
	height     int
}

func NewBuffer(path string) (*Buffer, error) {
	return nil, nil
}

func (b *Buffer) Insert() {
}

func (b *Buffer) Delete() {
}

func (b *Buffer) Render() {
}
