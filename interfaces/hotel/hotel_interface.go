package hotel_interface

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
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
