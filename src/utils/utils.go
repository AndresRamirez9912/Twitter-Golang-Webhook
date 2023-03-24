package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

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
