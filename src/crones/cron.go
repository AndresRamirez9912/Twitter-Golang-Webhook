package crones

import (
	"fmt"
	"log"
	"twitter-webhook/src/services"

	"github.com/robfig/cron" //Biblioteca que contiene las funciones del cron.
)

func Cron() {

	c := cron.New() //Creamos objeto de cron

	c.AddFunc("*/5 * * * * *", func() {

		_, err := services.LookDirectMessages()

		//Handle Error o manejador de errores
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Cron funcionando")
		//fmt.Printf("%+v\n", data) // Se imprime todo, la clave (identificador)

	})

	c.Start() //Se inicia el Cron

	defer c.Stop() //Se detiene el Cron

	select {} //Para que se ejecute infinitamente

}
