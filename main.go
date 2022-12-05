package main

import (
	"net/http"

	"go-webhook/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
