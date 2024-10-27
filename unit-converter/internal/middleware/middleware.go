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
