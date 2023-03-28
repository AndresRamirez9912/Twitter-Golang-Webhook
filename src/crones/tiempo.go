package crones

import (
	"fmt"
	"log"
	"time"
	"twitter-webhook/src/services"
)

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

	data, err := services.LookDirectMessages()

	//Handle Error o manejador de errores
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cron funcionando")
	fmt.Printf("%+v\n", data) // Se imprime todo, la clave (identificador)
}
