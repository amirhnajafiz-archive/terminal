package command

import (
	"fmt"
	"time"

	"github.com/amirhnajafiz/terminal/server/internal/types"
)

func NewCommand(u string, a func([]string) (string, error)) *types.Command {
	return &types.Command{
		Use:    u,
		Action: a,
	}
}

func timer(_ []string) (string, error) {
	return fmt.Sprintln(time.Now()), nil
}

func Register() []*types.Command {
	return []*types.Command{
		NewCommand("time", timer),
	}
}
