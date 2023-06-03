package term

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/jsniffen/geditor/gui"
)

func hideCursor(b *bytes.Buffer) {
	fmt.Fprint(b, "\033[?25l")
}

func moveCursor(b *bytes.Buffer, x, y int) {
	fmt.Fprintf(b, "\033[%d;%dH", x, y)
}

func setForeground(b *bytes.Buffer, c gui.Color) {
	fmt.Fprintf(b, "\033[38;2;%d;%d;%dm", c.Red, c.Green, c.Blue)
}

func setBackground(b *bytes.Buffer, c gui.Color) {
	fmt.Fprintf(b, "\033[48;2;%d;%d;%dm", c.Red, c.Green, c.Blue)
}

func getEvent() (Event, error) {
	var e Event

	rd := bufio.NewReader(os.Stdin)

	buf := make([]byte, 4)
	n, err := rd.Read(buf)
	if err != nil {
		return e, err
	}

	if n == 1 {
		if buf[0] == 127 {
			e.KeyCode = KeyBackspace
		} else {
			e.KeyCode = buf[0]
		}
	} else if n == 3 {
		if buf[0] == '\033' && buf[1] == '[' {
			if buf[2] == 'A' {
				e.KeyCode = KeyUp
			} else if buf[2] == 'B' {
				e.KeyCode = KeyDown
			} else if buf[2] == 'C' {
				e.KeyCode = KeyRight
			} else if buf[2] == 'D' {
				e.KeyCode = KeyLeft
			}
		}
	} else if n == 4 {
		if buf[0] == '\033' && buf[1] == '[' && buf[3] == '~' {
			if buf[2] == '3' {
				e.KeyCode = KeyDelete
			}
		}
	}

	return e, nil
}

func GetEvents() chan Event {
	c := make(chan Event)

	go func() {
		for {
			ev, err := getEvent()
			if err == nil {
				c <- ev
			}
		}
	}()

	return c
}

func Render(cells []gui.Cell) {
	b := bytes.Buffer{}

	hideCursor(&b)
	moveCursor(&b, 1, 1)

	bg := gui.Color{}
	setBackground(&b, bg)

	fg := gui.Color{}
	setForeground(&b, fg)

	for _, cell := range cells {
		if cell.Background != bg {
			bg = cell.Background
			setBackground(&b, bg)
		}

		if cell.Foreground != fg {
			fg = cell.Foreground
			setForeground(&b, fg)
		}

		r := utf8.AppendRune(nil, cell.Rune)
		b.Write(r)
	}
	b.WriteTo(os.Stdout)
}
