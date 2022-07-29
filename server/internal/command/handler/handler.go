package handler

import (
	"fmt"
	"time"
)

func GetTime(_ []string) (string, error) {
	return fmt.Sprintln(time.Now()), nil
}
