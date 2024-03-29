package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/models"
	"twitter-webhook/src/oauth"
)

// Write direct message to one person
func SendDirectMesage(userId string, message string) error {
	urlSendMessage := fmt.Sprintf(constants.SEND_DIRECT_MESSAGE_ENDPOINT, userId)
	oaut := oauth.CreateoAuth(urlSendMessage, constants.POST)

	dMBody := &models.DirectMessageRequestBody{
		Text: message,
	}
	jsonBody, err := json.Marshal(dMBody)
	if err != nil {
		log.Fatal(err)
		return err
	}

	body, err := oaut.SendOAuthRequest(jsonBody)
	if err != nil {
		log.Fatal(err)
		return err
	}

	dmResponse := &models.DirectMessageResponse{}
	err = json.Unmarshal(body, dmResponse)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Get all direct messages
func LookDirectMessages() (*models.LookUpDirectMessageResponse, error) {
	// Add query parameters
	baseURL, err := url.Parse(constants.LOOKUP_DIRECT_MESSAGES)
	if err != nil {
		log.Fatal(err)
	}
	queryParams := url.Values{}
	queryParams.Set(constants.DM_EVENT_FIELDS_QUERY, constants.DM_EVENT_FIELDS_VALUE)
	baseURL.RawQuery = queryParams.Encode()
	urlLookUp, err := url.QueryUnescape(baseURL.String())
	if err != nil {
		log.Fatal(err)
	}
	oaut := oauth.CreateoAuth(urlLookUp, constants.GET)
	body, err := oaut.SendOAuthRequest(nil)
	if err != nil {
		log.Fatal(err)
	}

	lookUpDMResponse := &models.LookUpDirectMessageResponse{}
	err = json.Unmarshal(body, lookUpDMResponse)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return lookUpDMResponse, nil
}

// Get direct messages by Id
func LookDirectMessagesById(id string) (*models.LookUpDirectMessageResponse, error) {
	// Add query parameters
	baseURL, err := url.Parse(fmt.Sprintf(constants.LOOKUP_DIRECT_MESSAGES_BY_ID, id))
	if err != nil {
		log.Fatal(err)
	}
	queryParams := url.Values{}
	queryParams.Set(constants.DM_EVENT_FIELDS_QUERY, constants.DM_EVENT_FIELDS_VALUE)
	baseURL.RawQuery = queryParams.Encode()
	URL, err := url.QueryUnescape(baseURL.String())
	if err != nil {
		log.Fatal(err)
	}
	oaut := oauth.CreateoAuth(URL, constants.GET)
	body, err := oaut.SendOAuthRequest(nil)
	if err != nil {
		log.Fatal("Error at the moment to send the request", err)
	}
	lookUpDMResponse := &models.LookUpDirectMessageResponse{}
	err = json.Unmarshal(body, lookUpDMResponse)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return lookUpDMResponse, nil
}
