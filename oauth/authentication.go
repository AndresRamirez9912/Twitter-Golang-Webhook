package oauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type oauth struct {
	consumer_key        string
	consumer_key_secret string
	base_URL            string
	method              string
}

func CreateoAuth(Consumer_key string, Consumer_key_secret string, Base_URL string, Method string) *oauth {
	return &oauth{
		consumer_key:        Consumer_key,
		consumer_key_secret: Consumer_key_secret,
		base_URL:            Base_URL,
		method:              Method,
	}
}

func (auth oauth) GetoauthToken() {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := uuid.NewString()
	req, err := auth.createRequest(timestamp, nonce)
	if err != nil {

	}
	auth.createBaseString(timestamp, nonce)
	fmt.Println("Request formada::::", req.URL.String())
}

func (auth oauth) createRequest(timestamp string, nonce string) (*http.Request, error) {
	req, err := http.NewRequest(auth.method, auth.base_URL, nil)
	if err != nil {
		log.Print("Error Creating the request")
		return nil, err
	}
	query := req.URL.Query()
	query.Add("oauth_consumer_key", auth.consumer_key)
	query.Add("oauth_signature_method", "HMAC-SHA1")
	query.Add("oauth_timestamp", timestamp)
	query.Add("oauth_nonce", nonce)
	query.Add("oauth_version", "1.0")
	query.Add("oauth_signature", "")

	req.URL.RawQuery = query.Encode()
	return req, nil
}

func createSignature(baseURL string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(baseURL))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (auth oauth) createBaseString(timestamp string, nonce string) {
	parameters := map[string]string{
		"oauth_consumer_key":     auth.consumer_key,
		"oauth_signature_method": auth.method,
		"oauth_timestamp":        timestamp,
		"oauth_nonce":            nonce,
		"oauth_version":          "1.0",
	}

	orderedKeys := make([]string, 0, len(parameters))
	for k := range parameters {
		orderedKeys = append(orderedKeys, k)
	}
	sort.Strings(orderedKeys)
	encodedParameters := " "
	for k := range orderedKeys {
		if encodedParameters == " " {
			encodedParameters = fmt.Sprintf("%s=%s", orderedKeys[k], parameters[orderedKeys[k]])
		} else {
			encodedParameters += encodedParameters + fmt.Sprintf("&%s=%s", orderedKeys[k], parameters[orderedKeys[k]])
		}
	}
	fmt.Println("Base Stirng::::", encodedParameters)
}
