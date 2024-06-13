package user_middlewares

import (
	"encoding/json"
	user_interface "hotel-booking-golang-gin/interfaces/user"
	error_handler "hotel-booking-golang-gin/util/error"
	response_handler "hotel-booking-golang-gin/util/response"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserValidator(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		c.Abort()
		return
	}
	var bodyContent map[string]interface{}

	err = json.Unmarshal(bodyBytes, &bodyContent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing JSON"})
		c.Abort()
		return
	}
	if bodyContent["role"] == nil {
		bodyContent["role"] = "USER"
	}

	user_obj := user_interface.User{
		Name:     bodyContent["name"].(string),
		Email:    bodyContent["email"].(string),
		Password: bodyContent["password"].(string),
		Role:     bodyContent["role"].(string),
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(&user_obj)

	if err != nil {
		c.JSON(http.StatusBadRequest, response_handler.Error("VALIDATION_ERROR", &error_handler.ErrArg{
			Code:        "VALIDATION_ERROR",
			Description: err.Error(),
		}))
		c.Abort()
		return
	}

	c.Set("User", user_obj)
	c.Next()
}

func UserLoginValidator(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		c.Abort()
		return
	}
	var bodyContent map[string]interface{}

	err = json.Unmarshal(bodyBytes, &bodyContent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing JSON"})
		c.Abort()
		return
	}

	user_obj := user_interface.UserLogin{
		Email:    bodyContent["email"].(string),
		Password: bodyContent["password"].(string),
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(&user_obj)

	if err != nil {
		c.JSON(http.StatusBadRequest, response_handler.Error("VALIDATION_ERROR", &error_handler.ErrArg{
			Code:        "VALIDATION_ERROR",
			Description: err.Error(),
		}))
		c.Abort()
		return
	}

	c.Set("User", user_obj)
	c.Next()
}
