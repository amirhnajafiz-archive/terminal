package command

import "github.com/amirhnajafiz/terminal/server/internal/types"

func NewCommand(u string, a func([]string) error) *types.Command {
	return &types.Command{
		Use:    u,
		Action: a,
	}
}
