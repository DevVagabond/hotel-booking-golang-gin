package hotel_middleware

import (
	"encoding/json"
	"fmt"
	"hotel-booking-golang-gin/initializers"
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

func ValidateBooking(context *gin.Context) {
	user := context.MustGet("User").(user_interface.User)

	bodyBytes, error := io.ReadAll(context.Request.Body)

	if error != nil {
		context.JSON(400, gin.H{"error": "error"})
		context.Abort()
		return
	}

	bodyContent := hotel_interface.BookingInput{}

	json.Unmarshal(bodyBytes, &bodyContent)

	validator := validator.New(validator.WithRequiredStructEnabled())

	error = validator.Struct(&bodyContent)

	if error != nil {
		context.JSON(400, gin.H{"error": error.Error()})
		context.Abort()
		return
	}

	bodyContent.UserID = user.ID
	bodyContent.IsPaid = true

	//function to check room availability

	room := hotel_interface.HotelRoom{}

	initializers.DB.First(&room, bodyContent.RoomID)

	if room.ID == 0 {
		context.JSON(400, gin.H{"error": "Invalid room id"})
		context.Abort()
		return
	}

	var overlappingBookings int64
	initializers.DB.Model(&hotel_interface.Booking{}).Where(
		"room_id = ? AND ((check_in <= ? AND check_out >= ?) OR (check_in < ? AND check_out >= ?) OR (check_in <= ? AND check_out > ?))",
		bodyContent.RoomID, bodyContent.CheckIn, bodyContent.CheckIn, bodyContent.CheckOut, bodyContent.CheckOut, bodyContent.CheckIn, bodyContent.CheckOut,
	).Count(&overlappingBookings)

	if overlappingBookings > 0 {
		context.JSON(400, gin.H{"error": "Room is not available for the selected dates"})
		context.Abort()
		return
	}

	stayDuration := bodyContent.CheckOut.Sub(bodyContent.CheckIn).Hours() / 24

	totalPrice := stayDuration * float64(room.RentPrice)

	bodyContent.TotalCost = float32(totalPrice)

	context.Set("Booking", bodyContent)
	context.Next()

}
