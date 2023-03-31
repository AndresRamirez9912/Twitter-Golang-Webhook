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
	"os"
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
	URL string,
	method string) *oauth {

	return &oauth{
		consumer_key:        os.Getenv(constants.API_KEY),        // API KEY,
		consumer_key_secret: os.Getenv(constants.API_KEY_SECRET), // API Secret
		URL:                 constants.BASE_URL + URL,
		method:              method,
		oauth_token:         os.Getenv(constants.ACCESS_TOKEN),        // Access Token
		oauth_token_secret:  os.Getenv(constants.ACCESS_TOKEN_SECRET), // Access Secret
	}
}

func (auth oauth) SendOAuthRequest(body []byte, queries map[string]string) ([]byte, error) {
	req, err := auth.CreateOAuthRequest(body, queries)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	bodyResponse, err := utils.SendRequest(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return bodyResponse, nil
}

func (auth oauth) CreateOAuthRequest(body []byte, queries map[string]string) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))

	baseString := auth.createBaseString(timestamp, nonce, queries)
	signature := createSignature(baseString, auth.consumer_key_secret, auth.oauth_token_secret)
	req, err := auth.createRequest(timestamp, nonce, signature, body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return req, nil
}

func (auth oauth) createBaseString(timestamp string, nonce string, queries map[string]string) string {
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
	encodedQueries := ""

	for k, v := range queries {
		if encodedQueries == "" {
			encodedQueries = url.QueryEscape(fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
		} else {
			encodedQueries += url.QueryEscape(fmt.Sprintf("&%s=%s", k, url.QueryEscape(v)))
		}
	}

	encodedParameters := ""
	for k := range orderedKeys {
		if encodedParameters == "" && encodedQueries == "" {
			encodedParameters = url.QueryEscape(fmt.Sprintf("%s=%s", orderedKeys[k], parameters[orderedKeys[k]]))
		} else {
			encodedParameters += url.QueryEscape(fmt.Sprintf("&%s=%s", orderedKeys[k], parameters[orderedKeys[k]]))
		}
	}
	return fmt.Sprintf("%s&%s&%s", strings.ToUpper(auth.method), url.QueryEscape(auth.URL), (encodedQueries + encodedParameters))
}

func createSignature(baseURL string, consumer_secret string, oauth_token_secret string) string {
	signing_key := fmt.Sprintf("%s&%s", url.QueryEscape(consumer_secret), url.QueryEscape(oauth_token_secret))
	h := hmac.New(sha1.New, []byte(signing_key))
	h.Write([]byte(baseURL))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (auth oauth) createRequest(timestamp string, nonce string, signature string, body []byte) (*http.Request, error) {

	req, err := http.NewRequest(auth.method, auth.URL+"?"+constants.DM_EVENT_FIELDS_QUERY+"="+constants.DM_EVENT_FIELDS_VALUE, bytes.NewBuffer(body))
	if err != nil {
		log.Print("Error Creating the request")
		return nil, err
	}

	authHeader := fmt.Sprintf(constants.AUTHORIZATION_TEMPLATE, auth.consumer_key, nonce, url.QueryEscape(signature), timestamp, auth.oauth_token)
	req.Header = http.Header{
		constants.ACCEPT:        {"*/*"},
		constants.AUTHORIZATION: {authHeader},
		constants.CONNECTION:    {"close"},
	}
	return req, nil
}

func (auth oauth) GetValidationLink(oauthParameters map[string]string) {
	fmt.Printf(constants.VALIDATION_LINK_TEMPLATE+"\n", oauthParameters[constants.OAUTH_TOKEN])
}
