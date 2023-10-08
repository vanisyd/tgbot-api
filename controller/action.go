package controller

import (
	database "github.com/vanisyd/tgbot-db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetActions(userID primitive.ObjectID) []database.Action {
	return database.FindActions(userID)
}
