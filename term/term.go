package term

import "os"
import "fmt"

func HideCursor() {
	os.Stdout.Write([]byte("\033[?25l"))
}

func ShowCursor() {
	os.Stdout.Write([]byte("\033[?25h"))
}

func PrintColor(r, g, b uint8) {
	s := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	os.Stdout.Write([]byte(s))
}
