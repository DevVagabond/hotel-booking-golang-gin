package hotel_service

import (
	"hotel-booking-golang-gin/initializers"
	hotel_interface "hotel-booking-golang-gin/interfaces/hotel"
	error_handler "hotel-booking-golang-gin/util/error"
)

func CreateHotel(hotel hotel_interface.HotelInput) (hotel_interface.Hotel, *error_handler.ErrArg) {
	hotelModel := hotel_interface.Hotel{
		Name:       hotel.Name,
		Address:    hotel.Address,
		Phone:      hotel.Phone,
		Email:      hotel.Email,
		Website:    hotel.Website,
		IsActive:   hotel.IsActive,
		IsVerified: false,
		Latitude:   hotel.Latitude,
		Longitude:  hotel.Longitude,
		OwnerID:    hotel.OwnerID,
	}

	// Save the hotel

	hotelResponse := hotel_interface.Hotel{}
	initializers.DB.Create(&hotelModel)
	initializers.DB.First(&hotelResponse, hotelModel.ID)

	if hotelResponse.ID == 0 {
		return hotel_interface.Hotel{}, &error_handler.ErrArg{
			Code:        "HOTEL_CREATION_FAILED",
			Description: "Hotel creation failed",
			Title:       "Hotel creation failed",
		}
	}

	return hotelResponse, nil
}
