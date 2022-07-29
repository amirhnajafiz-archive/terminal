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

func Register(h handler.Handler) []*types.Command {
	return []*types.Command{
		create("time", h.GetTime),
		create("whoami", h.GetUser),
		create("os", h.GetOS),
	}
}
