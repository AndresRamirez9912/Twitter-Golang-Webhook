package handler

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello From Golang Webhook")
}

func TestHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello From Golang Webhook")
}
