package middlewares

import (
	"net/http"
)

func AdminProtection(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		// tmp, err := template.ParseFiles("../web/html/admin/signin.html")
		// if err != nil {
		// 	http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		// }

		// err = tmp.Execute(res, nil)
		// if err != nil {
		// 	http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		// }
		// return
		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}
