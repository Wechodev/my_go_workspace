package main

import (
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("."))
	http.ListenAndServeTLS(":8001","cert.pem", "key.pem", h)
}
