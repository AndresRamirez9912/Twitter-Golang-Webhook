package oauth

import (
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

func (auth oauth) SendOAuthRequest(body []byte) ([]byte, error) {
	req, err := auth.createOAuthRequest(body)
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

func (auth oauth) createOAuthRequest(body []byte) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))

	req, err := utils.CreateRequest(auth.method, auth.URL, body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	baseString := auth.createBaseString(timestamp, nonce, req)
	signature := createSignature(baseString, auth.consumer_key_secret, auth.oauth_token_secret)
	auth.sign(signature, timestamp, nonce, req)
	return req, nil
}

func (auth oauth) createBaseString(timestamp string, nonce string, req *http.Request) string {
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

	for k, v := range req.URL.Query() {
		if encodedQueries == "" {
			encodedQueries = url.QueryEscape(fmt.Sprintf("%s=%s", k, url.QueryEscape(v[0])))
		} else {
			encodedQueries += url.QueryEscape(fmt.Sprintf("&%s=%s", k, url.QueryEscape(v[0])))
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
	return fmt.Sprintf("%s&%s&%s", strings.ToUpper(auth.method), url.QueryEscape(constants.BASE_URL+req.URL.Path), (encodedQueries + encodedParameters))
}

func createSignature(baseURL string, consumer_secret string, oauth_token_secret string) string {
	signing_key := fmt.Sprintf("%s&%s", url.QueryEscape(consumer_secret), url.QueryEscape(oauth_token_secret))
	h := hmac.New(sha1.New, []byte(signing_key))
	h.Write([]byte(baseURL))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (auth oauth) sign(signature string, timestamp string, nonce string, req *http.Request) {
	authHeader := fmt.Sprintf(constants.AUTHORIZATION_TEMPLATE, auth.consumer_key, nonce, url.QueryEscape(signature), timestamp, auth.oauth_token)
	req.Header.Add(constants.AUTHORIZATION, authHeader)
}

func (auth oauth) GetValidationLink(oauthParameters map[string]string) {
	fmt.Printf(constants.VALIDATION_LINK_TEMPLATE+"\n", oauthParameters[constants.OAUTH_TOKEN])
}
