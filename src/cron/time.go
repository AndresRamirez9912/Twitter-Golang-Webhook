package cron

import (
	"log"
	"time"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/database"
	"twitter-webhook/src/models"
	"twitter-webhook/src/services"
)

var Cont = 0

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
	//Handle Error o manejador de errores
	if err != nil {
		log.Fatal(err)
	}
	// Comparar si son diferentes (iterar la cantidad de mensajes nuevos)
	if Cont != messages.Meta.Result_count && messages.Data[0].Sender_id != constants.ID_BOT {
		entry := models.TwitterField{
			Id:            messages.Data[0].Sender_id,
			CreatedAt:     time.Now().String(),
			Active:        true,
			LastMessageId: messages.Data[0].Id,
		}
		err := database.CreateItem(entry)
		if err != nil {
			log.Fatal(err)
		}
		Cont = messages.Meta.Result_count
		// Enviar Respuesta al cliente
		services.SendDirectMesage(entry.Id, "El bot del banco Popular te da la Bienvenida")
	}
}
