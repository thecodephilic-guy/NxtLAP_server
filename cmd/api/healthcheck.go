package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthcheckHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	})
}
