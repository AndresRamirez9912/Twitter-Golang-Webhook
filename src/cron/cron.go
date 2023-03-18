package cron

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {

	c := cron.New()

	defer c.Stop()

	c.AddFunc("* * * * *", func() {
		fmt.Println("Verificando la hora por segundo :", time.Now().String())

	})

	c.Start()

	select {}
}
