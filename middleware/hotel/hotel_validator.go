package hotel_middleware

import (
	"encoding/json"
	"fmt"
	hotel_interface "hotel-booking-golang-gin/interfaces/hotel"
	user_interface "hotel-booking-golang-gin/interfaces/user"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HotelValidator(context *gin.Context) {
	// Validate the incoming request

	fmt.Print("HotelValidator middleware\n", context.MustGet("User"))
	owner := context.MustGet("User").(user_interface.User)
	requestBodyBytes, error := io.ReadAll(context.Request.Body)

	if error != nil {
		context.JSON(400, gin.H{"error": "error"})
		context.Abort()
		return
	}

	var requestBodyContent = make(map[string]interface{})

	error = json.Unmarshal(requestBodyBytes, &requestBodyContent)

	if error != nil {
		context.JSON(400, gin.H{"error": "error parsing JSON"})
		context.Abort()
		return
	}

	// Set default values
	if requestBodyContent["IsActive"] == nil {
		requestBodyContent["IsActive"] = true
	}

	if requestBodyContent["website"] == nil {
		requestBodyContent["website"] = ""
	}
	if requestBodyContent["Latitude"] == nil {
		requestBodyContent["Latitude"] = 0.0
	}
	if requestBodyContent["Longitude"] == nil {
		requestBodyContent["Longitude"] = 0.0
	}

	fmt.Println("=== lon====", reflect.TypeOf(requestBodyContent["Longitude"]))
	fmt.Println("=== lat====", requestBodyContent["Latitude"])

	hotelObj := hotel_interface.HotelInput{
		Name:       requestBodyContent["name"].(string),
		Address:    requestBodyContent["address"].(string),
		Phone:      requestBodyContent["phone"].(string),
		Email:      requestBodyContent["email"].(string),
		Website:    requestBodyContent["website"].(string),
		IsActive:   requestBodyContent["IsActive"].(bool),
		IsVerified: false,
		Latitude:   float32(requestBodyContent["Latitude"].(float64)),
		Longitude:  float32(requestBodyContent["Longitude"].(float64)),
		OwnerID:    owner.ID,
	}

	validator := validator.New(validator.WithRequiredStructEnabled())

	error = validator.Struct(&hotelObj)

	if error != nil {
		context.JSON(400, gin.H{"error": error.Error()})
		context.Abort()
		return
	}

	context.Set("Hotel", hotelObj)
	context.Next()
}
