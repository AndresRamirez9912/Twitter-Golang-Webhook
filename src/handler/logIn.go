package handler

import (
	"net/http"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/oauth"
	"twitter-webhook/src/utils"
)

func LogIn(w http.ResponseWriter, req *http.Request) {
	oaut := oauth.CreateoAuth(
		constants.REQUEST_TOKEN_URL, // Base URL
		constants.POST)              // Access Secret
	oauthParameters := utils.GetOAuthParameters(string(oaut.SendOAuthRequest()))
	oaut.GetValidationLink(oauthParameters)
}
