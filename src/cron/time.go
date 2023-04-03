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

var LastMessage = "1642990609106452485"

func WaitMessages() chan bool {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				report()
			case <-quit:
				ticker.Stop()
			}
		}
	}()
	return quit
}

func report() {

	messages, err := services.LookDirectMessages()
	if err != nil {
		log.Fatal(err)
	}
	msgCount := 0
	for messages.Data[msgCount].Id != LastMessage {
		if constants.ID_BOT != messages.Data[msgCount].Sender_id {
			entry := models.TwitterField{
				Id:            messages.Data[msgCount].Sender_id,
				CreatedAt:     time.Now().String(),
				Active:        true,
				LastMessageId: messages.Data[msgCount].Id,
			}
			err := database.CreateItem(entry)
			if err != nil {
				log.Fatal(err)
			}
			response := menu.Response(strings.ToLower(messages.Data[msgCount].Text))
			services.SendDirectMesage(entry.Id, response)
		}
		msgCount++
	}
	LastMessage = messages.Data[0].Id
}
