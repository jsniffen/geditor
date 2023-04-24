package term

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/jsniffen/geditor/gui"
)

func hideCursor(buffer []byte) {
	os.Stdout.Write([]byte("\033[?25l"))
}

func setForeground(b *bytes.Buffer, c gui.Color) {
	fmt.Fprintf(b, "\033[38;2;%d;%d;%dm", c.Red, c.Green, c.Blue)
}

func setBackground(b *bytes.Buffer, c gui.Color) {
	fmt.Fprintf(b, "\033[48;2;%d;%d;%dm", c.Red, c.Green, c.Blue)
}

func getEvent() (gui.Event, error) {
	var e gui.Event

	rd := bufio.NewReader(os.Stdin)

	buf := make([]byte, 4)
	n, err := rd.Read(buf)
	if err != nil {
		return e, err
	}

	if n == 1 {
		e.KeyCode = buf[0]
	}

	return e, nil
}

func GetEvents() chan gui.Event {
	c := make(chan gui.Event)

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
