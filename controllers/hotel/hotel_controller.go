package hotel_controller

import (
	hotel_interface "hotel-booking-golang-gin/interfaces/hotel"
	user_interface "hotel-booking-golang-gin/interfaces/user"
	hotel_service "hotel-booking-golang-gin/service/hotel"
	response_handler "hotel-booking-golang-gin/util/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateHotel(context *gin.Context) {

	hotel := context.MustGet("Hotel").(hotel_interface.HotelInput)

	hotelObj, error := hotel_service.CreateHotel(hotel)

	if error != nil {
		context.JSON(http.StatusBadRequest, response_handler.Error("HOTEL_CREATION_FAILED", error))
	} else {
		response := response_handler.OK(hotelObj)
		context.JSON(http.StatusOK, response)
	}
}

func ListHotel(context *gin.Context) {
	// Call the service function
	hotels := hotel_service.ListHotel(&hotel_interface.HotelQuery{})

	context.JSON(http.StatusOK, response_handler.OK(hotels))
}

func UpdateHotel(context *gin.Context) {
	user := context.MustGet("User").(user_interface.User)

	hotelUpdateData := context.MustGet("Hotel").(hotel_interface.HotelInput)

	hotelId := context.Param("id")

	hotelIdUint, err := strconv.ParseUint(hotelId, 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid hotel id"})
		context.Abort()
		return
	}

	hotelObj, error := hotel_service.UpdateHotel(hotelUpdateData, uint(hotelIdUint), user.ID)

	if error != nil {
		context.JSON(http.StatusBadRequest, response_handler.Error("HOTEL_UPDATE_FAILED", error))
	} else {
		response := response_handler.OK(hotelObj)
		context.JSON(http.StatusOK, response)
	}
}

func VerifyHotel(context *gin.Context) {

	hotelId := context.Param("id")

	hotelIdUint, err := strconv.ParseUint(hotelId, 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid hotel id"})
		context.Abort()
		return
	}

	hotelObj, error := hotel_service.VerifyHotel(uint(hotelIdUint))

	if error != nil {
		context.JSON(http.StatusBadRequest, response_handler.Error("HOTEL_VERIFICATION_FAILED", error))
	} else {
		response := response_handler.OK(hotelObj)
		context.JSON(http.StatusOK, response)
	}
}

func AddHotelRoom(context *gin.Context) {

	roomData := context.MustGet("Room").(hotel_interface.HotelRoomInput)
	dbresponse, error := hotel_service.AddHotelRoom(roomData)
	if error != nil {
		context.JSON(http.StatusBadRequest, response_handler.Error("HOTEL_VERIFICATION_FAILED", error))
	} else {
		response := response_handler.OK(dbresponse)
		context.JSON(http.StatusOK, response)
	}
}

func BookRoom(ctx *gin.Context) {
	bookingData := ctx.MustGet("Booking").(hotel_interface.BookingInput)

	res, error := hotel_service.BookRoom(bookingData)

	if error != nil {
		ctx.JSON(http.StatusBadRequest, response_handler.Error(error.Code, error))
	}

	ctx.JSON(http.StatusOK, res)

}
