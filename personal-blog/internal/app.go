package app

import (
	"log/slog"
	"net/http"
)

type Application struct {
	log *slog.Logger
}

func InitApplication(logger *slog.Logger) *Application {
	return &Application{
		log: logger,
	}
}

func (app *Application) ServerError(res http.ResponseWriter, err error) {
	app.log.Error("Internal server error", "error", err.Error())
	http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
