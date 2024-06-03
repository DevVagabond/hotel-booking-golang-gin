package main

import (
	"encoding/json"
	"hotel-booking-golang-gin/initializers"
	"io"
	"net/http"

	error_handler "hotel-booking-golang-gin/util/error"
	response_handler "hotel-booking-golang-gin/util/response"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
}

func main() {
	r := gin.Default()

	r.POST("/ping", func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
			return
		}
		var bodyContent map[string]interface{}

		err = json.Unmarshal(bodyBytes, &bodyContent)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing JSON"})
			return
		}

		if bodyContent["message"] == nil {
			error := error_handler.ErrArg{
				Title:       "Message is required",
				Code:        400,
				Description: "Message is required",
			}
			response := response_handler.Error(http.StatusBadRequest, &error)

			c.JSON(http.StatusBadRequest, response)
		} else {
			c.JSON(http.StatusOK, response_handler.OK(bodyContent))
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
