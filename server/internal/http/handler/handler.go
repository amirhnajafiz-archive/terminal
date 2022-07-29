package handler

import (
	"fmt"
	"net/http"

	"github.com/amirhnajafiz/terminal/server/internal/command"
	"github.com/amirhnajafiz/terminal/server/pkg/parser"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Commands map[string]*command.Command
}

func (h *Handler) Input(e echo.Context) error {
	var r Request

	_ = e.Bind(&r)

	out, err := h.commandHandler(r.Command)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	return e.String(http.StatusOK, out)
}

func (h *Handler) commandHandler(c string) (string, error) {
	parts, err := parser.ParseCommand(c)
	if err != nil {
		return "", err
	}

	if val, ok := h.Commands[parts[0]]; ok {
		rs, er := val.Action(nil)
		if er != nil {
			return "", er
		}

		return rs, nil
	}

	return "", fmt.Errorf("command not found: %s", parts[0])
}
