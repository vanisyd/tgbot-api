package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vanisyd/tgbot-api/controller"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	Router.GET("/options/weather", func(c *gin.Context) {
		userIDParam, ok := c.GetQuery("user_id")
		if ok {
			userId, err := primitive.ObjectIDFromHex(userIDParam)
			if err != nil {
				log.Fatal(err)
			}
			data := controller.GetActions(userId)
			if data != nil && len(data) > 0 {
				c.JSON(http.StatusOK, data)
			} else {
				c.JSON(http.StatusOK, gin.H{})
			}
		}
	})
}
