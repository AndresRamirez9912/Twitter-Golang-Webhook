package handler

import (
	"fmt"
	"net/http"
	"twitter-webhook/oauth"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello From Golang Webhook")
	fmt.Println("==========================================")
	oaut := oauth.Oauth{}
	oaut.Consumer_key = "WqPxHv9nXU3MdYhlrC2TXatBG"
	oaut.Consumer_key_secret = "cyxIJoqK82NOoMtS4SbzdHvskF4DP2rQp9WSC2tlesV1ZmwOW7"
	oaut.GetOauthToken()
}

func TestHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello From Golang Webhook")
}
