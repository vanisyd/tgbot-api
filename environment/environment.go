package environment

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Env EnvVars

func Init() {
	fmt.Println("Loading environment")
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load .env file")
	}

	Env.DBUri = os.Getenv("MONGODB_URI")
	Env.DBName = os.Getenv("MONGODB_DB_NAME")
}
