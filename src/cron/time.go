package cron

import (
	"log"
	"strings"
	"time"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/database"
	"twitter-webhook/src/menu"
	"twitter-webhook/src/models"
	"twitter-webhook/src/services"
)

var userRegister = &models.UsersActive{
	LastMessage: "",
	Users:       make(map[string]models.TwitterField),
}

func WaitMessages() chan bool {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				updateRegister()
				responseActve()
			case <-quit:
				ticker.Stop()
			}
		}
	}()
	return quit
}

func updateRegister() {
	messages, err := services.LookDirectMessages() // Get all messages
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages.Data { // Iterate the messages
		if message.Id == userRegister.LastMessage { // The message id must be new
			break // If there aren't new messages, stop iteration
		}

		if constants.ID_BOT == message.Sender_id { // The sender_Id must not be the Bot
			continue // If the message was sent by the bot, continue to the next message
		}

		err = updateUser(message.Sender_id) // The message is new and the sender_id isn't the bot
		if err != nil {
			log.Fatal(err)
		}

	}
	userRegister.LastMessage = messages.Data[0].Id
}

func responseActve() {
	for _, userId := range userRegister.ActiveUsers {
		// Get the messages by id
		messages, err := services.LookDirectMessagesById(userId)
		if err != nil {
			log.Fatal(err)
		}

		// Check If the sender was the bot or, the message is equal to the previous
		lastMessageID := userRegister.Users[userId].LastMessageId
		if (messages.Data[0].Sender_id == constants.ID_BOT) || (messages.Data[0].Id == lastMessageID) {
			continue // Continue to the next chat
		}

		// Send the response
		response := menu.Response(strings.ToLower(messages.Data[0].Text)) // Response the most recetly message
		err = services.SendDirectMesage(userId, response)
		if err != nil {

		}

		// Update the User in the Register
		userRegister.Users[userId] = models.TwitterField{
			Id:            userId,
			CreatedAt:     time.Now().String(),
			Active:        true,
			LastMessageId: messages.Data[0].Id,
		}

		// Update last message in the DB
		database.UpdateLastMessageId(userId, messages.Data[0].Id)
	}
}

func updateUser(senderId string) error {
	user, exist := userRegister.Users[senderId] // Check if exist the user into my register
	if exist && !user.Active {                  // User Exist but not Active
		if err := database.ChangeStatus(senderId, true); err != nil {
			log.Fatal(err)
			return err
		}
		user.Active = true
	} else if !exist { // User Not Exist
		entry := models.TwitterField{
			Id:        senderId,
			CreatedAt: time.Now().String(),
			Active:    true,
		}
		userRegister.ActiveUsers = append(userRegister.ActiveUsers, senderId)
		userRegister.Users[senderId] = entry
		if err := database.CreateItem(entry); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
