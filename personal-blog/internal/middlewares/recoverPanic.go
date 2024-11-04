package middlewares

import (
	"fmt"
	"net/http"
)

func (mw *MiddleWare) RecoverPanic(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				mw.app.ServerError(res, fmt.Errorf("%v", r))
			}
		}()

		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}
