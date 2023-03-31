package crones

import (
	"fmt"
	"log"
	"time"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/database"
	"twitter-webhook/src/models"
	"twitter-webhook/src/services"

	"github.com/google/uuid"
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

	fmt.Println("Cron Ejecutado:::", Cont)
	messages, err := services.LookDirectMessages()
	fmt.Println("Mensajes:::", messages.Meta.Result_count)
	//Handle Error o manejador de errores
	if err != nil {
		log.Fatal(err)
	}

	// Comparar si son diferentes (iterar la cantidad de mensajes nuevos)
	if Cont != messages.Meta.Result_count && messages.Data[0].Sender_id != constants.ID_BOT {
		entry := models.TwitterField{
			Id:            uuid.New().String(),
			CreatedAt:     time.Now().String(),
			ClientId:      uuid.New().String(),
			Active:        true,
			LastMessageId: messages.Data[0].Id,
		}
		err := database.CreateItem(entry)
		if err != nil {
			log.Fatal(err)
		}
		Cont = messages.Meta.Result_count
		// Enviar Respuesta al cliente
		services.SendDirectMesage("1638652102212743168", "El bot del banco Popular te da la Bienvenida")
	}
}
