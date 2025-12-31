package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/one-corp/NxtLAP_server/structs"
)

const baseUrl = "https://api.openf1.org/v1"

func (app *application) f1Handler(c *gin.Context) {
	year := c.DefaultQuery("year", strconv.Itoa(time.Now().Year()))

	response, err := http.Get(baseUrl + "/sessions?year=" + year)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status: %s", response.Status)
	}

	var data []structs.Session

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	c.IndentedJSON(http.StatusOK, data)
}
