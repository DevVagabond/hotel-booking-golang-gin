package user_interface

import (
	hotel_interface "hotel-booking-golang-gin/interfaces/hotel"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string                  `json:"name" validate:"required"`
	Email      string                  `json:"email" validate:"required,email"`
	Password   string                  `json:"password" validate:"required"`
	Role       string                  `json:"role"`
	IsActive   bool                    `json:"is_active"`
	IsVerified bool                    `json:"is_verified"`
	Hotels     []hotel_interface.Hotel `gorm:"foreignKey:OwnerID"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserSessionData struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     int64
}

type UserSession struct {
	gorm.Model
	UserID       uint   `json:"user_id" validate:"required"`
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
	ExpireAt     int64  `json:"expire_at" validate:"required"`
	Role         string `json:"role" validate:"required"`
}

type UserClaims struct {
	UserID uint
	Role   string
	jwt.StandardClaims
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireAt     int64  `json:"expire_at"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Role         string `json:"role"`
}
