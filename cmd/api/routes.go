package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	// Initialize a new Gin router instance.
	// gin.New() returns a blank slate.
	// If you wanted default middleware (logger/recovery), you would use gin.Default().
	router := gin.Default()

	//creating a group for v1 routes
	v1 := router.Group("/api/v1")
	{
		// Register the relevant methods, URL patterns and handler functions.
		// Gin provides semantic helper methods for each HTTP verb (GET, POST, etc).
		v1.GET("/healthcheck", app.healthcheckHandler)
		events := v1.Group("/events")
		{
			events.GET("/f1", app.f1Handler)
		}
	}

	// Return the router instance.
	return router
}
