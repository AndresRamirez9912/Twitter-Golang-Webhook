package utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"twitter-webhook/src/constants"
)

func CreateRequest(method string, URL string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(body))
	if err != nil {
		log.Print("Error Creating the request")
		return nil, err
	}

	req.Header = http.Header{
		constants.ACCEPT:     {"*/*"},
		constants.CONNECTION: {"close"},
	}
	return req, nil
}

func SendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetOAuthParameters(response string) map[string]string {
	result := make(map[string]string)
	data := strings.Split(response, "&")
	for _, parameter := range data {
		elements := strings.Split(parameter, "=")
		result[elements[0]] = elements[1]
	}
	return result
}
