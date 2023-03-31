package handler

import (
	"log"
	"net/http"
	"os"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/utils"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	oauth_token := r.URL.Query().Get(constants.OAUTH_TOKEN)
	oauth_verifier := r.URL.Query().Get(constants.OAUTH_VERIFIER)

	req, err := utils.CreateRequest(constants.POST, constants.BASE_URL+constants.ACCESS_TOKEN_URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add(constants.OAUTH_TOKEN, oauth_token)
	q.Add(constants.OAUTH_VERIFIER, oauth_verifier)

	req.URL.RawQuery = q.Encode()

	body, err := utils.SendRequest(req)
	if err != nil {
		log.Fatal(err)
	}
	// Store the Credentials into env variables
	accessTokens := utils.GetOAuthParameters(string(body))
	os.Setenv(constants.OAUTH_TOKEN, accessTokens[constants.OAUTH_TOKEN])
	os.Setenv(constants.OAUTH_TOKEN_SECRET, accessTokens[constants.OAUTH_TOKEN_SECRET])
	os.Setenv(constants.SCREEN_NAME, accessTokens[constants.SCREEN_NAME])
	os.Setenv(constants.USER_ID, accessTokens[constants.USER_ID])
}
