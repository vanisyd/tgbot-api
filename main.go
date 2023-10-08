package main

import (
	"github.com/vanisyd/tgbot-api/router"
	"log"
)

func main() {
	router.Init()

	if err := router.Router.Run(); err != nil {
		log.Fatal(err)
	}
}
