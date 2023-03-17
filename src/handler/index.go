package handler

import (
	"fmt"
	"net/http"
	"twitter-webhook/oauth"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello From Golang Webhook")
	oaut := oauth.CreateoAuth("WqPxHv9nXU3MdYhlrC2TXatBG", "cyxIJoqK82NOoMtS4SbzdHvskF4DP2rQp9WSC2tlesV1ZmwOW7", "https://api.twitter.com/oauth/request_token", "POST")
	oaut.GetoauthToken()
}

func TestHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello From Golang Webhook")
}
