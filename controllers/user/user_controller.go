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

func LoginUser(c *gin.Context) {

	// Call the service function

	user, error := user_service.LoginUser(c.MustGet("User").(user_interface.UserLogin))

	if error != nil {
		c.JSON(http.StatusBadRequest, response_handler.Error("USER_NOT_FOUND", error))
	} else {

		session, error := user_service.CreateSession(user)

		if error != nil {
			c.JSON(http.StatusBadRequest, response_handler.Error("SESSION_CREATION_ERROR", error))
			return
		}
		response := response_handler.OK(user_interface.UserLoginResponse{
			AccessToken:  session.AccessToken,
			RefreshToken: session.RefreshToken,
			ExpireAt:     session.ExpireAt,
			Name:         user.Name,
			Email:        user.Email,
			Role:         user.Role,
		})
		c.JSON(http.StatusOK, response)
	}
}
