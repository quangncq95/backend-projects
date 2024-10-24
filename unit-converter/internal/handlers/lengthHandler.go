package handler

import (
	"html/template"
	"log"
	"net/http"
)

func LengthHandlerGet(res http.ResponseWriter, req *http.Request) {
	ts, err := template.ParseFiles("../web/template/base.html")

	if err != nil {
		log.Fatalf("Error : %v", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(res, nil)
	if err != nil {
		log.Fatalf("Error :%v", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func LengthHandlerPost(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world"))
}
