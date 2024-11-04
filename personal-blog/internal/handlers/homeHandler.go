package handlers

import "net/http"

func (handler *Handler) HomeHandler(res http.ResponseWriter, req *http.Request) {
	panic("oops! something went wrong")
	res.Write([]byte("Hello world"))
}
