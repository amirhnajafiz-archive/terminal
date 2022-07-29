package terminal

import (
	"fmt"
	"time"
)

type Terminal struct {
	OS   string
	User string
}

func (_ *Terminal) GetHelp(_ []string) (string, error) {
	s := `Welcome to terminal.
	[help]   get terminal help
	[time]   get current time
	[whoami] get current user
	[os]     get user operating system
`

	return s, nil
}

func (_ *Terminal) GetTime(_ []string) (string, error) {
	return fmt.Sprintln(time.Now()), nil
}

func (t *Terminal) GetUser(_ []string) (string, error) {
	return t.User, nil
}

func (t *Terminal) GetOS(_ []string) (string, error) {
	return t.OS, nil
}
