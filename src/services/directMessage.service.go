package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/models"
	"twitter-webhook/src/oauth"
	"twitter-webhook/src/utils"
)

func SendDirectMesage(userId string, message string) {
	urlSendMessage := fmt.Sprintf(constants.SEND_DIRECT_MESSAGE_ENDPOINT, userId)
	oaut := oauth.CreateoAuth(
		os.Getenv(constants.API_KEY),             // API KEY
		os.Getenv(constants.API_KEY_SECRET),      // API Secret
		urlSendMessage,                           // URL with the id of the user to send message
		constants.POST,                           // Method
		os.Getenv(constants.ACCESS_TOKEN),        // Access Token
		os.Getenv(constants.ACCESS_TOKEN_SECRET)) // Access Secret

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
	fmt.Println(dmResponse)
}

func LookDirectMessages() {
	oaut := oauth.CreateoAuth(
		os.Getenv(constants.API_KEY),             // API KEY
		os.Getenv(constants.API_KEY_SECRET),      // API Secret
		constants.LOOKUP_DIRECT_MESSAGES,         // URL with the id of the user to send message
		constants.GET,                            // Method
		os.Getenv(constants.ACCESS_TOKEN),        // Access Token
		os.Getenv(constants.ACCESS_TOKEN_SECRET)) // Access Secret

	req, err := oaut.CreateOAuthRequest(nil)
	if err != nil {
		log.Fatal(err)
	}

	body, err := utils.SendRequest(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
