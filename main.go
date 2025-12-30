package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/one-corp/NxtLAP_server/structs"
)

func main() {

	router := gin.Default()
	router.GET("/api/sessions", func(ctx *gin.Context) {
		year := ctx.Query("year")
		response, err := http.Get(F1_API_URL + "/sessions?year=" + year)

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

		ctx.IndentedJSON(http.StatusOK, data)
	})

	router.Run()

}
