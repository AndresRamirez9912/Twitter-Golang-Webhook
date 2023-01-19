package handler

import (
	"fmt"
	"net/http"
)

func WebhookHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Webhook path")
}
