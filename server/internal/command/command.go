package command

import (
	"github.com/amirhnajafiz/terminal/server/internal/command/handler"
	"github.com/amirhnajafiz/terminal/server/internal/types"
)

func NewCommand(u string, a func([]string) (string, error)) *types.Command {
	return &types.Command{
		Use:    u,
		Action: a,
	}
}

func Register() []*types.Command {
	return []*types.Command{
		NewCommand("time", handler.GetTime),
	}
}
