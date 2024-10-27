package main

import (
	"log"
	"log/slog"
	handler "ncquang/unit-converter/internal/handlers"
	"ncquang/unit-converter/internal/middleware"
	"ncquang/unit-converter/internal/router"
	"net/http"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	handler := handler.NewAppHandler(logger)
	logger.Info("Server start at port 8000\n")

	err := http.ListenAndServe(":8000", middleware.LogRequest(logger, router.Routes(handler)))
	log.Fatal("Error", err)
}
