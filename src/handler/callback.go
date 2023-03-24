package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/oauth"
)

func Authorize(w http.ResponseWriter, req *http.Request) {
	oauth_token := req.URL.Query().Get(constants.OAUTH_TOKEN)
	oauth_verifier := req.URL.Query().Get(constants.OAUTH_VERIFIER)

	req, err := createRequest(oauth_token, oauth_verifier)
	if err != nil {

	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

	}
	accessTokens := oauth.GetOAuthParameters(string(body))
	fmt.Println(accessTokens)
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
	fmt.Println(req.URL.String())
	return req, nil
}
