package hotel_controller

import (
	hotel_interface "hotel-booking-golang-gin/interfaces/hotel"
	hotel_service "hotel-booking-golang-gin/service/hotel"
	response_handler "hotel-booking-golang-gin/util/response"
	"net/http"

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
}
