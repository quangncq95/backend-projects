package handler

import "net/http"

func WeightHandlerGet(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world"))
}

func WeightHandlerPost(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world"))
}
