package handlers

import (
	"log/slog"
	app "ncquang/personal-blog/internal"
)

type Handler struct {
	log *slog.Logger
}

func NewHandler(logger *slog.Logger, app *app.Application) *Handler {
	return &Handler{
		log: logger,
	}
}
