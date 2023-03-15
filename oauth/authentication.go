package oauth

import (
	"fmt"
	"log"
	"net/http"
)

type Oauth struct {
	Consumer_key        string
	Consumer_key_secret string
}

func (auth Oauth) GetOauthToken() {
	req, err := auth.createRequest()
	if err != nil {

	}
	fmt.Println(req.URL.String())
}

func (auth Oauth) createRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", "https://api.twitter.com/oauth/authorize", nil)
	if err != nil {
		log.Print("Error Creating the request")
		return nil, err
	}
	query := req.URL.Query()
	query.Add("oauth_consumer_key", auth.Consumer_key)
	query.Add("oauth_signature_method", "HMAC-SHA1")
	query.Add("oauth_timestamp", "")
	query.Add("oauth_nonce", "")
	query.Add("oauth_version", "1.0")
	query.Add("oauth_signature", "")

	req.URL.RawQuery = query.Encode()
	return req, nil
}
