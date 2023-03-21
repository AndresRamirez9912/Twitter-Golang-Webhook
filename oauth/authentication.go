package oauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type oauth struct {
	consumer_key        string
	consumer_key_secret string
	base_URL            string
	method              string
	oauth_token         string
	oauth_token_secret  string
}

func CreateoAuth(
	api_key string,
	api_key_secret string,
	base_URL string,
	method string,
	access_token string,
	access_secret string) *oauth {

	return &oauth{
		consumer_key:        api_key,
		consumer_key_secret: api_key_secret,
		base_URL:            base_URL,
		method:              method,
		oauth_token:         access_token,
		oauth_token_secret:  access_secret,
	}
}

func (auth oauth) GetoauthToken() {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))

	baseString := auth.createBaseString(timestamp, nonce)
	signature := createSignature(baseString, auth.consumer_key_secret, auth.oauth_token_secret)

	req, err := auth.createRequest(timestamp, nonce, signature)
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
	fmt.Println(string(body))
}

func (auth oauth) createRequest(timestamp string, nonce string, signature string) (*http.Request, error) {
	req, err := http.NewRequest(auth.method, auth.base_URL, nil)
	if err != nil {
		log.Print("Error Creating the request")
		return nil, err
	}

	authHeade := fmt.Sprintf(`OAuth oauth_consumer_key="%s", oauth_nonce="%s", oauth_signature="%s", oauth_signature_method="HMAC-SHA1", oauth_timestamp="%s", oauth_token="%s", oauth_version="1.0"`, auth.consumer_key, nonce, url.QueryEscape(signature), timestamp, auth.oauth_token)

	// Add Headers
	req.Header = http.Header{
		"Accept":        {"*/*"},
		"Authorization": {authHeade},
		"Connection":    {"close"},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}
	return req, nil
}

func createSignature(baseURL string, consumer_secret string, oauth_token_secret string) string {
	signing_key := fmt.Sprintf("%s&%s", url.QueryEscape(consumer_secret), url.QueryEscape(oauth_token_secret))
	h := hmac.New(sha1.New, []byte(signing_key))
	h.Write([]byte(baseURL))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (auth oauth) createBaseString(timestamp string, nonce string) string {
	parameters := map[string]string{
		"oauth_consumer_key":     auth.consumer_key,
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        timestamp,
		"oauth_nonce":            nonce,
		"oauth_version":          "1.0",
		"oauth_token":            auth.oauth_token,
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
	return fmt.Sprintf("%s&%s&%s", strings.ToUpper(auth.method), url.QueryEscape(auth.base_URL), encodedParameters)
}
