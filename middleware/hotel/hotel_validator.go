package hotel_middleware

import (
	"encoding/json"
	"fmt"
	hotel_interface "hotel-booking-golang-gin/interfaces/hotel"
	user_interface "hotel-booking-golang-gin/interfaces/user"
	hotel_service "hotel-booking-golang-gin/service/hotel"
	"io"
	"reflect"
	"strconv"

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

func HotelUpdateValidator(context *gin.Context) {
	// Validate the incoming request
	hotelInput := hotel_interface.HotelInput{}

	requestBodyBytes, error := io.ReadAll(context.Request.Body)

	if error != nil {
		context.JSON(400, gin.H{"error": "error"})
		context.Abort()
		return
	}

	json.Unmarshal(requestBodyBytes, &hotelInput)

	validator := validator.New(validator.WithRequiredStructEnabled())

	error = validator.Struct(&hotelInput)

	if error != nil {
		context.JSON(400, gin.H{"error": error.Error()})
		context.Abort()
		return
	}

	hotelId := context.Param("id")
	if hotelId == "" {
		context.JSON(400, gin.H{"error": "Invalid hotel id"})
		context.Abort()
		return
	}

	hotelIdUint, err := strconv.ParseUint(hotelId, 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid hotel id"})
		context.Abort()
		return
	}

	context.Set("HotelId", uint(hotelIdUint))

	context.Set("Hotel", hotelInput)
	context.Next()

}

func ValidateHotelRoomInput(context *gin.Context) {
	bodyBytes, error := io.ReadAll(context.Request.Body)

	if error != nil {
		context.JSON(400, gin.H{"error": "error"})
		context.Abort()
		return
	}

	bodyContent := hotel_interface.HotelRoomInput{}

	json.Unmarshal(bodyBytes, &bodyContent)

	validator := validator.New(validator.WithRequiredStructEnabled())

	error = validator.Struct(&bodyContent)

	if error != nil {
		context.JSON(400, gin.H{"error": error.Error()})
		context.Abort()
		return
	}

	user := context.MustGet("User").(user_interface.User)

	hotelQuery := hotel_interface.HotelQuery{
		ID:      bodyContent.HotelID,
		OwnerID: user.ID,
	}

	hotel := hotel_service.ListHotel(&hotelQuery)

	if len(hotel) == 0 {
		context.JSON(400, gin.H{"error": "Invalid hotel id"})
		context.Abort()
		return
	}

	context.Set("Room", bodyContent)

	context.Next()
}
