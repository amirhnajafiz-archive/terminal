package command

import (
	"github.com/amirhnajafiz/terminal/server/internal/command/terminal"
)

type Command struct {
	Use    string                              `json:"use"`
	Action func(args []string) (string, error) `json:"action"`
}

func create(u string, a func([]string) (string, error)) *Command {
	return &Command{
		Use:    u,
		Action: a,
	}
}

func Register(t terminal.Terminal) map[string]*Command {
	return map[string]*Command{
		"help":   create("help", t.GetHelp),
		"time":   create("time", t.GetTime),
		"whoami": create("whoami", t.GetUser),
		"os":     create("os", t.GetOS),
	}
}
