package handler

import (
	"fmt"
	"net/http"
	"twitter-webhook/src/services"
)

func WebhookHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Webhook path")
	// services.SendDirectMesage("1638652102212743168", "Hello Buddy")
	services.LookDirectMessages()
}
