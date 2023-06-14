package term

import "golang.org/x/sys/unix"
import "os"

var originalTermios *unix.Termios

func Init() error {
	f, err := os.Open("/dev/tty")
	if err != nil {
		return err
	}
	defer f.Close()

	fd := int(f.Fd())

	originalTermios, err = unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return err
	}

	tios := &unix.Termios{}
	*tios = *originalTermios

	tios.Iflag &= ^(uint32(unix.IGNBRK) | uint32(unix.BRKINT) | uint32(unix.PARMRK) | uint32(unix.ISTRIP) | uint32(unix.INLCR) | uint32(unix.IGNCR) | uint32(unix.ICRNL) | uint32(unix.IXON))
        tios.Oflag &= ^uint32(unix.OPOST)
        tios.Lflag &= ^(uint32(unix.ECHO) | uint32(unix.ECHONL) | uint32(unix.ICANON) | uint32(unix.ISIG) | uint32(unix.IEXTEN))
        tios.Cflag &= ^(uint32(unix.CSIZE) | uint32(unix.PARENB))
        tios.Cflag |= uint32(unix.CS8)
        tios.Cc[unix.VMIN] = 1
        tios.Cc[unix.VTIME] = 0

	unix.IoctlSetTermios(fd, unix.TCSETSF, tios)

	return nil
}

func GetSize() (uint32, uint32, error) {
	ws, err := unix.IoctlGetWinsize(unix.Stdin, unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}

	return uint32(ws.Col), uint32(ws.Row), nil
}

func deInit() error {
	f, err := os.Open("/dev/tty")
	if err != nil {
		return err
	}
	defer f.Close()

	fd := int(f.Fd())
	unix.IoctlSetTermios(fd, unix.TCSETS, originalTermios)

	return nil
}
