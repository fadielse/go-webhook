package main

import (
	"net/http"

	"https://github.com/fadielse/go-webhook/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
