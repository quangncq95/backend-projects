package middleware

import (
	"log/slog"
	"net/http"
)

func LogRequest(logger *slog.Logger, next http.Handler) http.Handler {
	fun := func(res http.ResponseWriter, req *http.Request) {
		var (
			method = req.Method
			url    = req.URL.Path
		)
		logger.Info("received request", "Method", method, "URL", url)
		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fun)
}

func AddHeaders(nextHandler http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		header := res.Header()
		header.Set("Content-Security-Policy", "default-src 'self'; form-action 'self'; object-src 'none'; frame-ancestors 'none'; upgrade-insecure-requests; block-all-mixed-content")
		header.Set("X-Frame-Options", "deny")
		header.Set("Referrer-Policy", "no-referrer")
		header.Set("X-Content-Type-Options", "nosniff")
		nextHandler.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}
