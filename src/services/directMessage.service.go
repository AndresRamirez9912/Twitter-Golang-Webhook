package services

import (
	"encoding/json"
	"fmt"
	"log"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/models"
	"twitter-webhook/src/oauth"
	"twitter-webhook/src/utils"
)

// Write direct message to one person
func SendDirectMesage(userId string, message string) {
	urlSendMessage := fmt.Sprintf(constants.SEND_DIRECT_MESSAGE_ENDPOINT, userId)
	oaut := oauth.CreateoAuth(urlSendMessage, constants.POST)

	dMBody := &models.DirectMessageRequestBody{
		Text: message,
	}
	jsonBody, err := json.Marshal(dMBody)
	if err != nil {
		log.Fatal(err)
	}

	req, err := oaut.CreateOAuthRequest(jsonBody)
	if err != nil {
		log.Fatal(err)
	}

	body, err := utils.SendRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	dmResponse := &models.DirectMessageResponse{}
	err = json.Unmarshal(body, dmResponse)
	if err != nil {

	}
	fmt.Printf("%+v\n", dmResponse)
}

// Get all direct messages
// Los segundos parentesis son los datos de salida
func LookDirectMessages() (*models.LookUpDirectMessageResponse, error) {
	oaut := oauth.CreateoAuth(constants.LOOKUP_DIRECT_MESSAGES, constants.GET)

	req, err := oaut.CreateOAuthRequest(nil)
	if err != nil {
		log.Fatal(err)
	}

	body, err := utils.SendRequest(req)
	if err != nil {
		log.Fatal(err)
	}
	lookUpDMResponse := &models.LookUpDirectMessageResponse{}
	err = json.Unmarshal(body, lookUpDMResponse)
	if err != nil {
		log.Fatal(err)
	}
	return lookUpDMResponse, nil

}
