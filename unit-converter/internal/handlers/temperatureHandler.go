package handler

import "net/http"

func TemperatureHandlerGet(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world"))
}

func TemperatureHandlerPost(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world"))
}
