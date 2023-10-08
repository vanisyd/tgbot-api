package bootstrap

import (
	"github.com/vanisyd/tgbot-api/environment"
	"github.com/vanisyd/tgbot-api/router"
	database "github.com/vanisyd/tgbot-db"
	"log"
)

func Init() {
	environment.Init()
	database.Init(environment.Env.DBUri, environment.Env.DBName)
	router.Init()

	if err := router.Router.Run(); err != nil {
		log.Fatal(err)
	}
}
