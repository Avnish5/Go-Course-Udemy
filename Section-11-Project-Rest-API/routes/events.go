package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": "failed to featch events"})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not prase event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "failed to fetch event by id"})
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse requsted data"})
		return
	}

	event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "failed to create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not prase event id"})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "failed to fetch event by id"})
		fmt.Println(err)
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse requsted data"})
		fmt.Println(err)
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.UpdateEventById()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "failed to fetch event by id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event update successfully"})

}

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not prase event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "failed to fetch event by id"})
		return
	}

	err = event.DeleteEventById()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "failed to delete event by id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "delete event successfully"})

}
