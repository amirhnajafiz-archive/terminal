package handler

import (
	"fmt"
	"time"
)

type Handler struct {
	OS   string
	User string

	Current string
}

func (h *Handler) GetTime(_ []string) (string, error) {
	return fmt.Sprintln(time.Now()), nil
}

func (h *Handler) GetUser(_ []string) (string, error) {
	return h.User, nil
}

func (h *Handler) GetOS(_ []string) (string, error) {
	return h.OS, nil
}
