package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utkarsh-singh1/project/go/EventAPI/models"
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

	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest,fmt.Sprintln("Wrong Request Sent By User"))
		return
	}
	
	event.ID = 1

	c.JSON(http.StatusCreated, fmt.Sprintf("The current event is created and your registered event info is %v",event))
}
