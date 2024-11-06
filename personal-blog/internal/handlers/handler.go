package handlers

import (
	"log/slog"
	app "ncquang/personal-blog/internal"
)

type Handler struct {
	log *slog.Logger
	app *app.Application
}

func NewHandler(logger *slog.Logger, app *app.Application) *Handler {
	return &Handler{
		log: logger,
		app: app,
	}
}
