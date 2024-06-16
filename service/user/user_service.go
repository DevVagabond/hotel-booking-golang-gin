package user_service

import (
	"fmt"
	"hotel-booking-golang-gin/initializers"
	user_interface "hotel-booking-golang-gin/interfaces/user"
	crypto_util "hotel-booking-golang-gin/util/crypto"
	error_handler "hotel-booking-golang-gin/util/error"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateUser(userData user_interface.User) (user_interface.User, *error_handler.ErrArg) {
	existingUser := initializers.DB.Where(&user_interface.User{Email: userData.Email}).First(&user_interface.User{})
	fmt.Println("======", existingUser.RowsAffected)
	if existingUser.RowsAffected > 0 {
		return user_interface.User{}, &error_handler.ErrArg{
			Code:        "USER_ALREADY_EXISTS",
			Description: "User already exists",
			Title:       "User already exists",
		}
	}
	userData.Password = crypto_util.HashPassword(userData.Password)
	initializers.DB.Create(&userData)
	return userData, nil
}

func LoginUser(userData user_interface.UserLogin) (user_interface.User, *error_handler.ErrArg) {
	user := user_interface.User{}
	initializers.DB.Where(&user_interface.User{Email: userData.Email}).First(&user)
	fmt.Printf("%+v\n", user)
	if user.ID == 0 {
		return user_interface.User{}, &error_handler.ErrArg{
			Code:        "USER_NOT_FOUND",
			Description: "User not found",
			Title:       "User not found",
		}
	}
	if !crypto_util.CheckPasswordHash(userData.Password, user.Password) {
		return user_interface.User{}, &error_handler.ErrArg{
			Code:        "INVALID_PASSWORD",
			Description: "Invalid password",
			Title:       "Invalid password",
		}
	}

	return user, nil
}

func CreateSession(user user_interface.User) (user_interface.UserSessionData, *error_handler.ErrArg) {

	timeNow := time.Now()
	userClaim := user_interface.UserClaims{
		UserID: user.ID,
		Role:   user.Role,
	}
	accessToken, error := crypto_util.GenerateAccessToken(userClaim)

	if error != nil {
		return user_interface.UserSessionData{}, &error_handler.ErrArg{
			Code:        "ACCESStOKEN_GENERATION_ERROR",
			Description: "Error generating access token",
			Title:       "Error generating access token",
		}
	}

	jwtClaim := jwt.StandardClaims{
		IssuedAt:  timeNow.Unix(),
		ExpiresAt: timeNow.Add(time.Hour * 24).Unix(),
	}

	refreshToken, error := crypto_util.GenerateRefreshToken(jwtClaim)

	if error != nil {
		return user_interface.UserSessionData{}, &error_handler.ErrArg{
			Code:        "REFRESH_TOKEN_GENERATION_ERROR",
			Description: "Error generating refresh token",
			Title:       "Error generating refresh token",
		}
	}

	userSessionData := user_interface.UserSessionData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireAt:     jwtClaim.ExpiresAt,
	}

	userSessionObj := user_interface.UserSession{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireAt:     jwtClaim.ExpiresAt,
		Role:         user.Role,
	}

	initializers.DB.Create(&userSessionObj)

	return userSessionData, nil

}
