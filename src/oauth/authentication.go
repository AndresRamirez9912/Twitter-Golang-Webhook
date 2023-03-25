package oauth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/utils"

	"github.com/google/uuid"
)

type oauth struct {
	consumer_key        string
	consumer_key_secret string
	URL                 string
	method              string
	oauth_token         string
	oauth_token_secret  string
}

func CreateoAuth(
	api_key string,
	api_key_secret string,
	URL string,
	method string,
	access_token string,
	access_secret string) *oauth {

	return &oauth{
		consumer_key:        api_key,
		consumer_key_secret: api_key_secret,
		URL:                 constants.BASE_URL + URL,
		method:              method,
		oauth_token:         access_token,
		oauth_token_secret:  access_secret,
	}
}

func (auth oauth) SendOAuthRequest() []byte {
	req, err := auth.CreateOAuthRequest(nil) // I don't need Body
	if err != nil {
		log.Fatal(err)
	}
	body, err := utils.SendRequest(req)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (auth oauth) CreateOAuthRequest(body []byte) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))

	baseString := auth.createBaseString(timestamp, nonce)
	signature := createSignature(baseString, auth.consumer_key_secret, auth.oauth_token_secret)

	req, err := auth.createRequest(timestamp, nonce, signature, body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return req, nil
}

func (auth oauth) createBaseString(timestamp string, nonce string) string {
	parameters := map[string]string{
		constants.OAUTH_CONSUMER_KEY:     auth.consumer_key,
		constants.OAUTH_SIGNATURE_METHOD: constants.OAUTH_METHOD,
		constants.OAUTH_TIMESTAMP:        timestamp,
		constants.OAUTH_NONCE:            nonce,
		constants.OAUTH_VERSION:          "1.0",
		constants.OAUTH_TOKEN:            auth.oauth_token,
	}

	orderedKeys := make([]string, 0, len(parameters))
	for k := range parameters {
		orderedKeys = append(orderedKeys, k)
	}
	sort.Strings(orderedKeys)
	encodedParameters := " "
	for k := range orderedKeys {
		if encodedParameters == " " {
			encodedParameters = url.QueryEscape(fmt.Sprintf("%s=%s", orderedKeys[k], parameters[orderedKeys[k]]))
		} else {
			encodedParameters += url.QueryEscape(fmt.Sprintf("&%s=%s", orderedKeys[k], parameters[orderedKeys[k]]))
		}
	}
	return fmt.Sprintf("%s&%s&%s", strings.ToUpper(auth.method), url.QueryEscape(auth.URL), encodedParameters)
}

func createSignature(baseURL string, consumer_secret string, oauth_token_secret string) string {
	signing_key := fmt.Sprintf("%s&%s", url.QueryEscape(consumer_secret), url.QueryEscape(oauth_token_secret))
	h := hmac.New(sha1.New, []byte(signing_key))
	h.Write([]byte(baseURL))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (auth oauth) createRequest(timestamp string, nonce string, signature string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(auth.method, auth.URL, bytes.NewBuffer(body))
	if err != nil {
		log.Print("Error Creating the request")
		return nil, err
	}

	authHeade := fmt.Sprintf(constants.AUTHORIZATION_TEMPLATE, auth.consumer_key, nonce, url.QueryEscape(signature), timestamp, auth.oauth_token)

	req.Header = http.Header{
		constants.ACCEPT:        {"*/*"},
		constants.AUTHORIZATION: {authHeade},
		constants.CONNECTION:    {"close"},
		constants.CONTENT_TYPE:  {constants.APPLICATION_JSON},
	}
	return req, nil
}

func (auth oauth) GetValidationLink(oauthParameters map[string]string) {
	fmt.Printf(constants.VALIDATION_LINK_TEMPLATE, oauthParameters[constants.OAUTH_TOKEN])
}
