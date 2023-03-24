package handler

import (
	"net/http"
	"os"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/oauth"
)

func LogIn(w http.ResponseWriter, req *http.Request) {
	oaut := oauth.CreateoAuth(
		os.Getenv(constants.API_KEY),             // API KEY
		os.Getenv(constants.API_KEY_SECRET),      //API Secret
		constants.REQUEST_TOKEN_URL,              // Base URL
		constants.POST,                           // Method
		os.Getenv(constants.ACCESS_TOKEN),        // Access Token
		os.Getenv(constants.ACCESS_TOKEN_SECRET)) // Access Secret
	oauthParameters := oaut.GetoauthToken()
	oaut.GetValidationLink(oauthParameters)
}
