package handler

import (
	"net/http"

	"github.com/amirhnajafiz/terminal/server/internal/command"
	"github.com/amirhnajafiz/terminal/server/pkg/parser"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Commands map[string]*command.Command
}

func (h *Handler) Input(e echo.Context) error {
	var (
		r Request
	)

	_ = e.Bind(&r)

	parts, err := parser.ParseCommand(r.Command)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	if val, ok := h.Commands[parts[0]]; ok {
		rs, _ := val.Action(nil)

		return e.String(http.StatusOK, rs)
	}

	return e.String(http.StatusBadRequest, "empty request")
}
