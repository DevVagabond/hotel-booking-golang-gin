package user_controller

import (
	user_interface "hotel-booking-golang-gin/interfaces/user"
	user_service "hotel-booking-golang-gin/service/user"
	response_handler "hotel-booking-golang-gin/util/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	// Call the service function

	user, error := user_service.CreateUser(c.MustGet("User").(user_interface.User))

	if error != nil {
		c.JSON(http.StatusBadRequest, response_handler.Error("USER_ALREADY_EXISTS", error))
	} else {
		response := response_handler.OK(user)
		c.JSON(http.StatusOK, response)
	}
}
