package handler

import (
	"net/http"

	"github.com/amirhnajafiz/terminal/server/internal/types"
	"github.com/amirhnajafiz/terminal/server/pkg/parser"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Commands map[string]*types.Command
}

func (h *Handler) Input(e echo.Context) error {
	var (
		r  types.Request
		rs types.Response
	)

	_ = e.Bind(&r)

	parts, err := parser.ParseCommand(r.Command)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	if val, ok := h.Commands[parts[0]]; ok {
		rs.Result, err = val.Action(nil)

		return e.String(http.StatusOK, rs.Result)
	}

	return e.String(http.StatusBadRequest, "empty request")
}
