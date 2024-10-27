package router

import (
	handler "ncquang/unit-converter/internal/handlers"
	"net/http"
)

func Routes(appHandler *handler.AppHandler) http.Handler {
	mux := http.NewServeMux()

	fsHandler := http.FileServer(http.Dir("../web/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fsHandler))
	mux.HandleFunc("/{$}", appHandler.HomeHandlerGet)
	mux.HandleFunc("POST /result", appHandler.ConvertHandlerPost)

	return mux
}
