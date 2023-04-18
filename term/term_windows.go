package term

import "golang.org/x/sys/windows"

func Init() error {
	handle, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		return err
	}

	var mode uint32 
	err = windows.GetConsoleMode(handle, &mode)
	if err != nil {
		return err
	}

	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	err = windows.SetConsoleMode(handle, mode)
	if err != nil {
		return err
	}

	return nil
}
