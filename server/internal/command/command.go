package command

import (
	"github.com/amirhnajafiz/terminal/server/internal/command/terminal"
	"github.com/amirhnajafiz/terminal/server/internal/types"
)

func create(u string, a func([]string) (string, error)) *types.Command {
	return &types.Command{
		Use:    u,
		Action: a,
	}
}

func Register(t terminal.Terminal) map[string]*types.Command {
	return map[string]*types.Command{
		"time":   create("time", t.GetTime),
		"whoami": create("whoami", t.GetUser),
		"os":     create("os", t.GetOS),
	}
}
