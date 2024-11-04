package main

import (
	"log/slog"
	app "ncquang/personal-blog/internal"
	"ncquang/personal-blog/internal/handlers"
	"ncquang/personal-blog/internal/middlewares"
	"ncquang/personal-blog/internal/router"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := app.InitApplication(logger)
	mw := middlewares.NewMiddleWare(logger, app)
	handler := handlers.NewHandler(logger, app)
	mux := router.Routes(handler)
	http.ListenAndServe(":8000", mw.RecoverPanic(mw.CommonHeader(mw.LogRequest(mux))))
}
