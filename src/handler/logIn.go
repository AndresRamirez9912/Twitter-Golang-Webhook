package handler

import (
	"log"
	"net/http"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/oauth"
	"twitter-webhook/src/utils"
)

func LogIn(w http.ResponseWriter, req *http.Request) {
	oaut := oauth.CreateoAuth(
		constants.REQUEST_TOKEN_URL, // Base URL
		constants.POST)              // Access Secret
	body, err := oaut.SendOAuthRequest(nil)
	if err != nil {
		log.Fatal(err)
	}
	oauthParameters := utils.GetOAuthParameters(string(body))
	oaut.GetValidationLink(oauthParameters)
}
