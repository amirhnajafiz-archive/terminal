package cmd

import (
	"log"
	"os/user"
	"runtime"

	"github.com/amirhnajafiz/terminal/server/internal/command"
	"github.com/amirhnajafiz/terminal/server/internal/command/handler"
)

func Execute() {
	u, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	h := handler.Handler{
		User: u.Username,
		OS:   runtime.GOOS,
	}

	command.Register(h)
}
