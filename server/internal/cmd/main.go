package cmd

import (
	"log"
	"net/http"
	"os/user"
	"runtime"

	"github.com/amirhnajafiz/terminal/server/internal/command"
	"github.com/amirhnajafiz/terminal/server/internal/command/handler"
	"github.com/labstack/echo/v4"
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

	e := echo.New()
	e.GET("/api/cmd", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":5000"))
}
