package main

import (
	"fmt"
	"log"
	"ncquang/unit-converter/internal/router"
	"net/http"
)

func main() {
	fmt.Print("Server start at port 8000")
	err := http.ListenAndServe(":8000", router.Routes())
	log.Fatal("Error", err)
}
