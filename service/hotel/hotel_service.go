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

func ListHotel(query *hotel_interface.HotelQuery) []hotel_interface.Hotel {
	hotelQuery := hotel_interface.Hotel{
		IsActive:   true,
		IsVerified: true,
	}

	if (*query).ID != 0 {
		hotelQuery.ID = (*query).ID
	}
	if (*query).OwnerID != 0 {
		hotelQuery.OwnerID = (*query).OwnerID
	}

	hotels := []hotel_interface.Hotel{}
	initializers.DB.Where(hotelQuery).Find(&hotels)
	return hotels
}

func UpdateHotel(hotel hotel_interface.HotelInput, hotelId uint, owner uint) (hotel_interface.Hotel, *error_handler.ErrArg) {

	hotelObj := hotel_interface.Hotel{}
	initializers.DB.Where(&hotel_interface.Hotel{
		ID:      hotelId,
		OwnerID: owner,
	}).First(&hotelObj)

	if hotelObj.ID == 0 {
		return hotel_interface.Hotel{}, &error_handler.ErrArg{
			Code:        "HOTEL_NOT_FOUND",
			Description: "Hotel not found",
			Title:       "Hotel not found",
		}
	}

	hotelObj.Name = hotel.Name
	hotelObj.Address = hotel.Address
	hotelObj.Phone = hotel.Phone
	hotelObj.Email = hotel.Email
	hotelObj.Website = hotel.Website
	hotelObj.IsActive = hotel.IsActive
	hotelObj.Latitude = hotel.Latitude
	hotelObj.Longitude = hotel.Longitude

	res := initializers.DB.Save(&hotelObj)

	if res.Error != nil {
		return hotel_interface.Hotel{}, &error_handler.ErrArg{
			Code:        "HOTEL_UPDATE_FAILED",
			Description: "Hotel update failed",
			Title:       "Hotel update failed",
		}
	}

	return hotelObj, nil
}

func VerifyHotel(hotelId uint) (hotel_interface.Hotel, *error_handler.ErrArg) {
	hotelObj := hotel_interface.Hotel{}
	initializers.DB.Where(&hotel_interface.Hotel{
		ID: hotelId,
	}).First(&hotelObj)

	if hotelObj.ID == 0 {
		return hotel_interface.Hotel{}, &error_handler.ErrArg{
			Code:        "HOTEL_NOT_FOUND",
			Description: "Hotel not found",
			Title:       "Hotel not found",
		}
	}

	hotelObj.IsVerified = true

	res := initializers.DB.Save(&hotelObj)

	if res.Error != nil {
		return hotel_interface.Hotel{}, &error_handler.ErrArg{
			Code:        "HOTEL_VERIFICATION_FAILED",
			Description: "Hotel verification failed",
			Title:       "Hotel verification failed",
		}
	}

	return hotelObj, nil
}

func AddHotelRoom(room hotel_interface.HotelRoomInput) {
	hotelRoomObj := hotel_interface.HotelRoom{
		Name:        room.Name,
		Description: room.Description,
		RoomCount:   room.RoomCount,
		RentPrice:   room.RentPrice,
		HotelID:     room.HotelID,
	}

	hotelResponse := hotel_interface.HotelRoomResponse{}

	initializers.DB.Create(&hotelRoomObj)
	initializers.DB.First(&hotelResponse, hotelRoomObj.ID)

	for i := 0; i < len(room.AmenityList); i++ {
		initializers.DB.Create(&hotel_interface.RoomAmenities{
			AmenityType: room.AmenityList[i].AmenityType,
			Description: room.AmenityList[i].Description,
			HotelRoomID: hotelRoomObj.ID,
		})
	}

}
