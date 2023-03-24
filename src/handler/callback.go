package handler

import (
	"log"
	"net/http"
	"os"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/utils"
)

func Authorize(w http.ResponseWriter, req *http.Request) {
	oauth_token := req.URL.Query().Get(constants.OAUTH_TOKEN)
	oauth_verifier := req.URL.Query().Get(constants.OAUTH_VERIFIER)

	req, err := createRequest(oauth_token, oauth_verifier)
	if err != nil {

	}
	body, err := utils.SendRequest(req)
	if err != nil {

	}
	// Store the Credentials into env variables
	accessTokens := utils.GetOAuthParameters(string(body))
	os.Setenv(constants.OAUTH_TOKEN, accessTokens[constants.OAUTH_TOKEN])
	os.Setenv(constants.OAUTH_TOKEN_SECRET, accessTokens[constants.OAUTH_TOKEN_SECRET])
	os.Setenv(constants.SCREEN_NAME, accessTokens[constants.SCREEN_NAME])
	os.Setenv(constants.USER_ID, accessTokens[constants.USER_ID])
}

func createRequest(oauth_token string, oauth_verifier string) (*http.Request, error) {
	req, err := http.NewRequest(constants.POST, constants.ACCESS_TOKEN_URL, nil)
	if err != nil {
		log.Print("Error Creating the request")
		return nil, err
	}

	// Add the query parameters to valida
	q := req.URL.Query() // Copy the Query parameters
	q.Add(constants.OAUTH_TOKEN, oauth_token)
	q.Add(constants.OAUTH_VERIFIER, oauth_verifier)

	req.URL.RawQuery = q.Encode() // Assign the new query parameters into the request
	return req, nil
}
