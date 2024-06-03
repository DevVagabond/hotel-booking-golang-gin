package user_controller

import (
	"encoding/json"
	user_service "hotel-booking-golang-gin/service/user"
	response_handler "hotel-booking-golang-gin/util/response"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
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

	// Call the service function
	user := user_service.CreateUser(bodyContent)

	response := response_handler.OK(user)
	c.JSON(http.StatusOK, response)
}
