package handler

import "log/slog"

type AppHandler struct {
	logger *slog.Logger
}

func NewAppHandler(logger *slog.Logger) *AppHandler {
	return &AppHandler{
		logger,
	}
}
