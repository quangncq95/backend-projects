package middlewares

import "net/http"

func (mw *MiddleWare) CommonHeader(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("X-Frame-Options", "deny")
		// res.Header().Set("Content-Security-Policy", "default-src 'self'; form-action 'self'; object-src 'none'; frame-ancestors 'none'; upgrade-insecure-requests; block-all-mixed-content")
		res.Header().Set("Referrer-Policy", "no-referrer")
		res.Header().Set("X-Content-Type-Options", "nosniff")

		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}
