package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utkarsh-singh1/project/go/EventAPI/models"
	//"github.com/utkarsh-singh1/project/go/EventAPI/models"
)

func main() {

	server := gin.Default()

	server.GET("/events", getAllEvent)

	server.POST("/events", createNewEvent)
	
	
	server.Run(":8080")
}


func getAllEvent(c *gin.Context) {

	events := models.GetAllEvent()
	c.JSON(http.StatusOK, events)

}

func createNewEvent(c *gin.Context) {

	var event models.Event

	c.ShouldBindJSON(&event)

}
