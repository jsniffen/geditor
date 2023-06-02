package term

const (
	_ byte = iota
	KeyDelete
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyBackspace
)

type Event struct {
	KeyCode byte
}
