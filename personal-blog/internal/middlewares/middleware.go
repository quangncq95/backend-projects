package middlewares

import (
	"log/slog"
	app "ncquang/personal-blog/internal"
)

type MiddleWare struct {
	log *slog.Logger
	app *app.Application
}

func NewMiddleWare(logger *slog.Logger, app *app.Application) *MiddleWare {
	return &MiddleWare{
		log: logger,
		app: app,
	}
}
