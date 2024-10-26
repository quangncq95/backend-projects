package router

import (
	handler "ncquang/unit-converter/internal/handlers"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	fsHandler := http.FileServer(http.Dir("../web/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fsHandler))
	mux.HandleFunc("/{$}", handler.HomeHandlerGet)
	mux.HandleFunc("POST /result", handler.ConvertHandlerPost)

	return mux
}
