package crypto_util

import (
	user_interface "hotel-booking-golang-gin/interfaces/user"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		panic(error)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return error == nil
}

func GenerateAccessToken(userClaim user_interface.UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)

	token, error := accessToken.SignedString([]byte("secret"))

	return token, error
}

func GenerateRefreshToken(claims jwt.StandardClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, error := refreshToken.SignedString([]byte("refreshsecret"))

	return token, error
}
