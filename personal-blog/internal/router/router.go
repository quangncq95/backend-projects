package router

import (
	"ncquang/personal-blog/internal/handlers"
	"ncquang/personal-blog/internal/middlewares"
	"net/http"
)

func Routes(handler *handlers.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("../web/static"))))
	mux.HandleFunc("GET /{$}", handler.HomeHandler)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("GET /admin/{$}", handler.AdminHomeHandler)
	adminMux.HandleFunc("GET /admin/add", handler.AdminAddBlogHandlerGet)
	adminMux.HandleFunc("POST /upload-image", handler.UploadImage)

	mux.Handle("/", middlewares.AdminProtection(adminMux))

	return mux
}
