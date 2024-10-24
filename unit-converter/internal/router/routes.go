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
	mux.HandleFunc("POST /converter/length", handler.LengthHandlerPost)
	mux.HandleFunc("POST /converter/weight", handler.WeightHandlerPost)
	mux.HandleFunc("POST /converter/temperature", handler.TemperatureHandlerPost)

	return mux
}
