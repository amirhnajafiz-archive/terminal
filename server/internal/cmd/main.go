package cmd

import (
	"log"
	"os/user"
	"runtime"

	"github.com/amirhnajafiz/terminal/server/internal/command"
	"github.com/amirhnajafiz/terminal/server/internal/command/terminal"
	"github.com/amirhnajafiz/terminal/server/internal/http/handler"
	"github.com/labstack/echo/v4"
)

func Execute() {
	u, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	t := terminal.Terminal{
		User: u.Username,
		OS:   runtime.GOOS,
	}
	h := handler.Handler{
		Commands: command.Register(t),
	}
	e := echo.New()

	e.GET("/api/cmd", h.Input)

	e.Logger.Fatal(e.Start(":5000"))
}
