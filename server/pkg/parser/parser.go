package parser

import "strings"

const (
	separator = " "
)

func ParseCommand(command string) ([]string, error) {
	arg := strings.Split(command, separator)

	return arg, nil
}
