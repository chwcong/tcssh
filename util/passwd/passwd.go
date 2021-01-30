package passwd

import (
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
)

const (
	del       = 127
	backspace = 8
)

// GetPasswd get passwd from stdin and do not show the char to stdout
func GetPasswd() (pass []byte, err error) {
	stdinFd := int(os.Stdin.Fd())
	if terminal.IsTerminal(stdinFd) {
		if oldState, err := terminal.MakeRaw(stdinFd); err != nil {
			return pass, err
		} else {
			defer func() {
				err = terminal.Restore(stdinFd, oldState)
			}()
		}
	}
	Loop:
	for {
		char, err := getChar()
		if err != nil {
			break Loop
		}
		switch char {
		case '\n':
			break Loop
		case '\r':
			break Loop
		case del:
			length := len(pass)
			if length > 0 {
				pass = pass[:length-1]
			}
		case backspace:
			length := len(pass)
			if length > 0 {
				pass = pass[:length-1]
			}
		default:
			pass = append(pass, char)
		}
	}
	return pass, err
}

func getChar() (byte, error) {
	buf := make([]byte, 1)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		return 0, err
	}
	if n == 0 {
		return 0, io.EOF
	}
	return buf[0], nil
}
