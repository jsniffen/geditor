package edit

type GapBuffer struct {
	size   int
	start  int
	end    int
	buffer []byte
}

func NewGapBuffer(size int) *GapBuffer{
	b := make([]byte, size)
	return &GapBuffer{
		size: size,
		start: 0,
		end: size,
		buffer: b,
	}
}

func (gb *GapBuffer) String() string {
	return string(gb.buffer[0:gb.start]) + string(gb.buffer[gb.end:gb.size])
}

func (gb *GapBuffer) Insert(b byte) {
	gb.buffer[gb.start] = b
	gb.start += 1
}

func (gb *GapBuffer) Delete() {
	gb.start -= 1
}
