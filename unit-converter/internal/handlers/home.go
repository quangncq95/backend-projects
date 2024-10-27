package handler

import (
	"html/template"
	"log"
	"net/http"
)

func (appHandler *AppHandler) HomeHandlerGet(res http.ResponseWriter, req *http.Request) {
	ts, err := template.ParseFiles("../web/template/base.html", "../web/template/forms.html")

	if err != nil {
		log.Fatalf("Error : %v", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(res, "base", nil)
	if err != nil {
		log.Fatalf("Error :%v", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
