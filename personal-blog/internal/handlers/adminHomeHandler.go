package handlers

import (
	"net/http"
	"text/template"
)

func (handler *Handler) AdminHomeHandler(res http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("../web/html/base.html", "../web/html/admin/home.html")
	if err != nil {
		handler.app.ServerError(res, err)
	}

	err = tmp.ExecuteTemplate(res, "base", nil)
	if err != nil {
		handler.app.ServerError(res, err)
	}
}
