package hotel_interface

import (
	"time"

	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string      `json:"name" validate:"required"`
	Address    string      `json:"address" validate:"required"`
	Phone      string      `json:"phone" validate:"required"`
	Email      string      `json:"email" validate:"email"`
	Website    string      `json:"website"`
	IsActive   bool        `json:"is_active"`
	IsVerified bool        `json:"is_verified"`
	Latitude   float32     `json:"latitude"`
	Longitude  float32     `json:"longitude"`
	OwnerID    uint        `json:"owner_id"`
	Rooms      []HotelRoom `json:"rooms"`
}

type HotelInput struct {
	Name       string  `json:"name" validate:"required"`
	Address    string  `json:"address" validate:"required"`
	Phone      string  `json:"phone" validate:"required"`
	Email      string  `json:"email" validate:"email"`
	Website    string  `json:"website"`
	IsActive   bool    `json:"is_active"`
	IsVerified bool    `json:"is_verified"`
	Latitude   float32 `json:"latitude"`
	Longitude  float32 `json:"longitude"`
	OwnerID    uint    `json:"owner_id"`
}

type HotelResponse struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	Website    string  `json:"website"`
	IsActive   bool    `json:"is_active"`
	IsVerified bool    `json:"is_verified"`
	Latitude   float32 `json:"latitude"`
	Longitude  float32 `json:"longitude"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type HotelRoomInput struct {
	Name        string                    `json:"name" validate:"required"`
	Description string                    `json:"description" validate:"required"`
	AmenityList []HotelRoomAmenitiesInput `json:"amenityList"`
	RoomCount   int                       `json:"availableRoom" validate:"required"`
	RentPrice   float32                   `json:"rentPrice" validate:"required"`
	HotelID     uint                      `json:"hotelId" validate:"required"`
}

type HotelRoomAmenitiesInput struct {
	AmenityType string `json:"type" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type HotelQuery struct {
	ID         uint
	OwnerID    uint
	IsActive   bool
	IsVerified bool
}

type HotelRoom struct {
	gorm.Model
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description" validate:"required"`
	AmenityList []RoomAmenities `json:"amenityList"`
	RoomCount   int             `json:"availableRoom" validate:"required"`
	RentPrice   float32         `json:"rentPrice" validate:"required"`
	HotelID     uint            `json:"hotelId" validate:"required"`
}

type RoomAmenities struct {
	gorm.Model
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AmenityType string `json:"type" validate:"required"`
	Description string `json:"description" validate:"required"`
	HotelRoomID uint
}

type HotelRoomResponse struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description" validate:"required"`
	AmenityList []RoomAmenities `json:"amenityList"`
	RoomCount   int             `json:"availableRoom" validate:"required"`
	RentPrice   float32         `json:"rentPrice" validate:"required"`
	HotelID     uint            `json:"hotelId" validate:"required"`
}

type Booking struct {
	gorm.Model
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `json:"userId" gorm:"not null"`
	RoomID    uint      `json:"roomId" gorm:"not null"`
	CheckIn   time.Time `json:"checkIn" gorm:"not null"`
	CheckOut  time.Time `json:"checkOut" gorm:"not null"`
	TotalCost float32   `json:"totalCost" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"`
	IsPaid    bool      `json:"isPaid"`
}

type BookingInput struct {
	UserID    uint      `json:"userId" gorm:"not null"`
	RoomID    uint      `json:"roomId" gorm:"not null"`
	CheckIn   time.Time `json:"checkIn" gorm:"not null"`
	CheckOut  time.Time `json:"checkOut" gorm:"not null"`
	TotalCost float32   `json:"totalCost" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"`
	IsPaid    bool      `json:"isPaid"`
}
