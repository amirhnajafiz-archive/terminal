package command

import (
	"github.com/amirhnajafiz/terminal/server/internal/command/handler"
	"github.com/amirhnajafiz/terminal/server/internal/types"
)

func create(u string, a func([]string) (string, error)) *types.Command {
	return &types.Command{
		Use:    u,
		Action: a,
	}
}

func Register(h handler.Handler) map[string]*types.Command {
	return map[string]*types.Command{
		"time":   create("time", h.GetTime),
		"whoami": create("whoami", h.GetUser),
		"os":     create("os", h.GetOS),
	}
}
