package middlewares

import "net/http"

func (mw *MiddleWare) LogRequest(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		mw.log.Info("Received Request", "method", req.Method, "url", req.RequestURI)
		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}
