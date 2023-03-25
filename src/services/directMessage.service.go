package services

import (
	"fmt"
	"os"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/oauth"
	"twitter-webhook/src/utils"
)

func SendDirectMesage(userId string, message string) {
	urlSendMessage := constants.BASE_URL + fmt.Sprintf(constants.SEND_DIRECT_MESSAGE_ENDPOINT, userId)
	oaut := oauth.CreateoAuth(
		os.Getenv(constants.API_KEY),             // API KEY
		os.Getenv(constants.API_KEY_SECRET),      //API Secret
		urlSendMessage,                           // URL with the id of the user to send message
		constants.POST,                           // Method
		os.Getenv(constants.ACCESS_TOKEN),        // Access Token
		os.Getenv(constants.ACCESS_TOKEN_SECRET)) // Access Secret
	// Create Body

	// Create Request
	req, err := oaut.CreateOAuthRequest(nil)
	if err != nil {

	}

	// Send Request
	body, err := utils.SendRequest(req)
	if err != nil {

	}
	// Analyze response
	fmt.Println(string(body))
}
