package parser

import (
	"fmt"
	"strings"
)

const (
	separator = " "
)

func ParseCommand(command string) ([]string, error) {
	if len(command) == 0 {
		return nil, fmt.Errorf("empty command")
	}

	arg := strings.Split(command, separator)

	return arg, nil
}
