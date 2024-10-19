package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	autenticated := server.Group("/")
	autenticated.Use(middleware.Authneticate)
	autenticated.POST("/events", createEvent)
	autenticated.PUT("/events/:id", updateEvent)
	autenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)

}
