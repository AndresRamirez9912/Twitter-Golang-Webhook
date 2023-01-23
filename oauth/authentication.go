package oauth

import (
	"fmt"
	"net/http"

	"github.com/dghubble/oauth1"
)

var config = oauth1.Config{
	ConsumerKey:    "MY_CONSUMER_KEY",
	ConsumerSecret: "MY_CONSUMER_SECRET",
	CallbackURL:    "MY_CALL_BACK_URL",
}

func SetupOauth(w http.ResponseWriter, req *http.Request) {
	requestToken, _, err := config.RequestToken()
	if err != nil {
		fmt.Println("Error requesting the token")
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		fmt.Println("Error requesting the token")
	}
	http.Redirect(w, req, authorizationURL.String(), http.StatusFound)
}
