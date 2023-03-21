package handler

import (
	"fmt"
	"net/http"
	"twitter-webhook/oauth"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello From Golang Webhook")
	oaut := oauth.CreateoAuth(
		"G5XPaovsNIPYvIvBCIMEKumNv",                          // API KEY
		"7oymthawzdLEY8q3g1wgK6nfeLsRlKMmF0vhcy3CCPqMP1nBaX", //API Secret
		"https://api.twitter.com/oauth/request_token",        // Base URL
		"POST", // Method
		"579402887-Vmh4XvxiKGkt1ArgX7dvgapomn4UN6Ym20og2a9q", // Access Token
		"1QlIKAAS02MtyrJ12USWbNogRDndfiA9tBusa9ygN1E6r")      // Access Secret
	oaut.GetoauthToken()
}

func TestHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello From Golang Webhook")
}
