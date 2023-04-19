package term

import "golang.org/x/sys/windows"

func Init() error {
	var mode uint32

	in, err := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
	if err != nil {
		return err
	}

	err = windows.GetConsoleMode(in, &mode)
	if err != nil {
		return err
	}

	mode |= windows.ENABLE_VIRTUAL_TERMINAL_INPUT
	mode &= ^(uint32(windows.ENABLE_LINE_INPUT) | uint32(windows.ENABLE_ECHO_INPUT))

	err = windows.SetConsoleMode(in, mode)
	if err != nil {
		return err
	}

	out, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		return err
	}

	err = windows.GetConsoleMode(out, &mode)
	if err != nil {
		return err
	}

	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	err = windows.SetConsoleMode(out, mode)
	if err != nil {
		return err
	}

	return nil
}
