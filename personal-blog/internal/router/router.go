package router

import (
	"ncquang/personal-blog/internal/handlers"
	"net/http"
)

func Routes(handler *handlers.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", handler.HomeHandler)

	return mux
}
