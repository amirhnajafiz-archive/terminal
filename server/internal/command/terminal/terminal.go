package terminal

import (
	"fmt"
	"time"
)

type Terminal struct {
	OS   string
	User string

	Current string
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
