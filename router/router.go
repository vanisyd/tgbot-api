package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	Router.GET("/options/weather", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "TEST"})
	})
}
